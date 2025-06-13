package virtual_repo

import (
	"os"
	"path/filepath"

	"github.com/geektheripper/go-gutils/git/git_utils"
	"github.com/geektheripper/go-gutils/git/hack_ssh"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/jdx/go-netrc"
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

	if rurl.Protocol == "http" || rurl.Protocol == "https" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil
		}

		netrcPath := filepath.Join(homeDir, ".netrc")
		netrcFile, err := netrc.Parse(netrcPath)
		if err != nil {
			return nil
		}

		machine := netrcFile.Machine(rurl.Host)
		if machine != nil {
			v.auth = &http.BasicAuth{
				Username: machine.Get("login"),
				Password: machine.Get("password"),
			}
		}
	}

	return nil
}

func (v *VirtualRepo) GetAuthMethod() (transport.AuthMethod, error) {
	if v.auth != nil {
		err := v.EnsureAuth()
		if err != nil {
			return nil, err
		}
	}

	return v.auth, nil
}
