package virtual_repo

import (
	"os"

	"github.com/geektheripper/go-gutils/git/git_utils"
	"github.com/geektheripper/go-gutils/git/hack_ssh"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func (v *VirtualRepo) EnsureAuth() error {
	if v.auth != nil {
		return nil
	}

	rurl, _ := git_utils.ParseGitRemoteURL(v.remoteURL)

	if rurl.Protocol == "ssh" {
		if _, ok := os.LookupEnv("SSH_AUTH_SOCK"); ok {
			return nil
		}

		key, err := hack_ssh.GetKeyForHost(rurl.Host, rurl.User)
		if err != nil {
			return err
		}

		auth, err := ssh.NewPublicKeysFromFile(rurl.User, key, "")
		if err != nil {
			return err
		}

		v.auth = auth

		return nil
	}

	if rurl.Protocol == "http" {
		panic("http auth not implemented")
	}

	return nil
}
