package gitutils

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func GetContractGitCommit(rootDir string, ethurl string, contractAddress common.Address) ([]string,error){

	// checks if the rootDir is absolute else joins the path
	// provided with current working directory
	rootDirPath, err := filepath.Abs(rootDir)
	if err != nil {
		return nil,err
	}

	conn,err := ethclient.Dial(ethurl)
	if err != nil {
		return nil,err
	}

	RuntimeBytecode, err := conn.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return nil,err
	}

	r, err := git.PlainOpen(rootDirPath)
	if err != nil {
		return nil,err
	}

	ref, err := r.Head()
	if err != nil {
		return nil,err
	}


	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil,err
	}
	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		return nil,err
	}

	var gitCommits []string
	err = commitIter.ForEach(func(c *object.Commit) error {

		f, err := c.File("build/base_content_space.go")
		if err != object.ErrFileNotFound {
			content, _ := f.Contents()
			rBytecode := strings.Split(hexutil.Encode(RuntimeBytecode),"0x")[1]
			rBytecodePresent := strings.Contains(content, rBytecode)

			if rBytecodePresent {
				gitCommits = append(gitCommits, c.String())
			}
		}

		return nil
	})
	if err != nil {
		return nil,err
	}

	return gitCommits, nil
}

