package gitutils

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func GetContractGitCommit() error {

	address := "0x2f603b833d64fca173b4e4880b9f7935d1acc07f"
	clientUrl := "http://localhost:8545"

	conn,err := ethclient.Dial(clientUrl)
	if err != nil {
		return err
	}

	RuntimeBytecode, err := conn.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return err
	}

	projRootDir := "/Users/preethi/Desktop/qluvio/elv-io/contracts"
	r, err := git.PlainOpen(projRootDir)
	if err != nil {
		return err
	}

	ref, err := r.Head()
	if err != nil {
		return err
	}


	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}
	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		return err
	}

	err = commitIter.ForEach(func(c *object.Commit) error {
		hash := c.Hash.String()
		//line := strings.Split(c.Message, "\n")
		fmt.Println(hash)
		f, err := c.File("build/base_content_space.go")
		if err != object.ErrFileNotFound {
			content, _ := f.Contents()
			fmt.Println("CONTENT:", len(content))
			RuntimeBytecodePresent := strings.Contains(content, hex.EncodeToString(RuntimeBytecode))
			fmt.Println("RuntimeBytecode", RuntimeBytecodePresent)
			if RuntimeBytecodePresent {
				fmt.Println(commit.Hash, "bytehash is present")
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

