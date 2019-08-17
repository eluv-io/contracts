package gitutils

import (
"fmt"
"gopkg.in/src-d/go-git.v4"
"gopkg.in/src-d/go-git.v4/plumbing/object"
	"regexp"

)

func GetContractFile() error {

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
	//commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	//if err != nil {
	//	return err
	//}

	f, err := commit.File("build/base_content_space.go")
	if err != object.ErrFileNotFound {
		content, _ := f.Contents()

		re := regexp.MustCompile("([\"'`])(?:(?=(\\?))\2.)*?\1")
		fmt.Printf("Pattern: %v\n", re.String())      // print pattern
		fmt.Println("Matched:", re.MatchString(content)) // true

		fmt.Println("\nText between square brackets:")
		submatchall := re.FindAllString(content, -1)
		fmt.Println(submatchall)
		//for _, element := range submatchall {
		//	element = strings.Trim(element, "[")
		//	element = strings.Trim(element, "]")
		//	fmt.Println(element)
		//}


		//if strings.Contains(content, "BaseContentSpaceBin = ") {
		//	fmt.Println("FOUND")
		//}

	}

	//err = commitIter.ForEach(func(c *object.Commit) error {
	//	hash := c.Hash.String()
	//	line := strings.Split(c.Message, "\n")
	//	fmt.Println(hash, line)
	//	f, err := c.File("build/base_content_space.go")
	//	if err != object.ErrFileNotFound {
	//		content, _ := f.Contents()
	//		fmt.Println("CONTENT:",content)
	//
	//		if strings.Index(content, "BaseContentSpaceBin = ") != -1 {
	//			fmt.Println("FOUND")
	//		}
	//
	//	}
	//
	//	return nil
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}

