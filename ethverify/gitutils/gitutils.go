package gitutils

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func GetContractGitCommit(rootDir string, ethurl string, contractAddress common.Address) ([]string, error) {

	// checks if the rootDir is absolute else joins the path
	// provided with current working directory
	rootDirPath, err := filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}
	log.Debug("root directory", "path", rootDirPath)

	conn, err := ethclient.Dial(ethurl)
	if err != nil {
		return nil, err
	}

	RuntimeBytecode, err := conn.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return nil, err
	}
	if hexutil.Encode(RuntimeBytecode) == "0x" {
		return nil, fmt.Errorf("bytecode is not present in provided blockchain, url:%v", ethurl)
	}

	rBytecode := strings.Split(hexutil.Encode(RuntimeBytecode), "0x")[1]
	// retrieve swarm hash : 0xa1 0x65 'b' 'z' 'z' 'r' '0' 0x58 0x20 <32 bytes swarm hash> 0x00 0x29
	rBytecodeSwarmHash := rBytecode[len(rBytecode)-86:]
	log.Debug("swarm", "hash", rBytecodeSwarmHash)

	r, err := git.PlainOpen(rootDirPath)
	if err != nil {
		return nil, err
	}

	ref, err := r.Head()
	if err != nil {
		return nil, err
	}
	log.Debug("git HEAD", "value", ref.Hash().String())

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		return nil, err
	}

	var gitCommits []string
	err = commitIter.ForEach(func(c *object.Commit) error {
		log.Debug("commit iterated", "hash", c.Hash.String())

		f, err := c.File("build/base_content_space.go")
		if err != object.ErrFileNotFound {
			content, _ := f.Contents()

			rBytecodePresent := strings.Contains(content, rBytecodeSwarmHash)
			log.Debug("bytecode", "exists", rBytecodePresent)

			if rBytecodePresent {
				gitCommits = append(gitCommits, c.String())
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return gitCommits, nil
}
