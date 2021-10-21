package main

import (
	"fmt"
	"os"

	"github.com/n3wscott/releases-demo/semver/git"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage:\n semver <git url> <git-ref|version|branch>")
		os.Exit(1)
	}

	remote := os.Args[1]
	ref := os.Args[2]

	repo, err := git.GetRepo("HEAD", remote)
	if err != nil {
		panic(err)
	}

	next, err := repo.Next(ref)
	fmt.Printf("v%s", next)
}
