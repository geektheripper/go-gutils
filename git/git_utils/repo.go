package git_utils

import (
	"errors"
	"os"
	"path/filepath"
)

func IsGitRepo(dir string) (bool, error) {
	_, err := os.Stat(filepath.Join(dir, ".git", "HEAD"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func FindGitRoot(dir string) (string, error) {
	ok, err := IsGitRepo(dir)
	if err != nil {
		return "", err
	}

	if ok {
		return dir, nil
	}

	if dir == "." {
		dir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	parent := filepath.Dir(dir)
	if parent == dir {
		return "", errors.New("not in a git repo")
	}

	return FindGitRoot(parent)
}

type WorkingRepo struct {
	Root      string
	WdRel     string
	UnixWdRel string
}

func NewWorkingRepo() (*WorkingRepo, error) {
	root, err := FindGitRoot(".")
	if err != nil {
		return nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rel, err := filepath.Rel(root, wd)
	if err != nil {
		return nil, err
	}

	unixRel := filepath.ToSlash(rel)

	return &WorkingRepo{Root: root, WdRel: rel, UnixWdRel: unixRel}, nil
}
