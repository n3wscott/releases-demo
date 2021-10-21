package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/n3wscott/releases-demo/semver/git"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage:\n semver <git-ref|version|branch>")
		os.Exit(1)
	}
	ref := os.Args[1]

	fmt.Println("ref:", ref)
	base := filepath.Base(ref)
	fmt.Println("base:", base)
	version := base
	if strings.HasPrefix(version, "release-v") {
		version = version[len("release-v"):]
	}
	fmt.Println("version:", version)
	v, err := semver.Make(version)
	if err != nil {
		panic(err)
	}

	fmt.Println("major minor:", v.Major, v.Minor, v.Patch)

	repo, err := git.GetRepo("HEAD", "https://github.com/knative/eventing.git")
	if err != nil {
		panic(err)
	}
	fam := semver.Versions{}
	for _, tag := range repo.Tags {
		//fmt.Println("tag:", tag)
		if !strings.HasPrefix(tag, "v") {
			continue
		}
		tv, err := semver.Parse(tag[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse verison in branch: %s, %v", tag, err)
		}
		if tv.Major == v.Major && tv.Minor == v.Minor {
			fam = append(fam, tv)
		}
	}

	semver.Sort(fam)
	for _, f := range fam {
		fmt.Println(f.String())
	}

	next := fam[len(fam)-1]
	if err := next.IncrementPatch(); err != nil {
		panic(err)
	}
	fmt.Println("next:", next)
}
