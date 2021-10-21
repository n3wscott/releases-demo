package git

import (
	"fmt"
	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"path/filepath"
	"strings"
)

// Repo is a simplified git remote, containing only the list of tags, default
// branch and branches.
type Repo struct {
	Ref           string
	DefaultBranch string
	Tags          []string
	Branches      []string
}

// GetRepo will fetch a git repo and process it into a Repo object.
func GetRepo(ref, url string) (*Repo, error) {
	repo := new(Repo)
	repo.Ref = ref

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, ref := range refs {
		if ref.Name().IsTag() {
			repo.Tags = append(repo.Tags, ref.Name().Short())
		} else if ref.Name().IsBranch() {
			repo.Branches = append(repo.Branches, ref.Name().Short())
		} else if ref.Name() == "HEAD" { // Default branch.
			repo.DefaultBranch = ref.Target().Short()
		}
	}

	return repo, nil
}

// Next will parse ref (git ref, release branch, or version) and inspect
// the repo for tags that match the go mod release process to find the next
// patch version.
func (r *Repo) Next(ref string) (*semver.Version, error) {
	base := filepath.Base(ref)
	version := base
	if strings.HasPrefix(version, "release-") {
		version = version[len("release-"):]
	}
	v, err := semver.ParseTolerant(version)
	if err != nil {
		return nil, err
	}
	v.Patch = 0
	v.Pre = nil
	v.Build = nil

	fam := semver.Versions{}
	for _, tag := range r.Tags {
		if !strings.HasPrefix(tag, "v") {
			continue
		}
		tv, err := semver.ParseTolerant(tag[1:])
		if err != nil {
			return nil, fmt.Errorf("failed to parse verison in branch: %s, %v", tag, err)
		}
		if tv.Major == v.Major && tv.Minor == v.Minor {
			fam = append(fam, tv)
		}
	}

	semver.Sort(fam)

	if len(fam) > 0 {
		next := fam[len(fam)-1]
		if err := next.IncrementPatch(); err != nil {
			return nil, err
		}
		return &next, nil
	} else {
		return &v, nil
	}
}
