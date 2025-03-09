package git_utils

import (
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strings"
)

// https://git-scm.com/docs/git-clone#_git_urls

// URL format:
// ssh://[<user>@]<host>[:<port>]/<path-to-git-repo>
// git://<host>[:<port>]/<path-to-git-repo>
// http[s]://<host>[:<port>]/<path-to-git-repo>
// ftp[s]://<host>[:<port>]/<path-to-git-repo>

// SCP format:
// [<user>@]<host>:/<path-to-git-repo>

var GitRemoteScpRegex = regexp.MustCompile(`^((?P<user>git)@)?(?P<host>[a-zA-Z0-9-_\.]+)\:(?P<path>.*)$`)
var AvailableGitRemoteProtocols = []string{"ssh", "git", "http", "https", "ftp", "ftps"}

type GitRemote struct {
	URL      string
	Protocol string
	User     string
	Host     string
	Port     string
	Path     string
}

func ParseGitRemoteURL(remoteURL string) (*GitRemote, error) {
	if !strings.HasSuffix(remoteURL, ".git") {
		return nil, fmt.Errorf("remote URL must end with .git: %s", remoteURL)
	}

	if gurl, err := url.Parse(remoteURL); err == nil {
		if !slices.Contains(AvailableGitRemoteProtocols, gurl.Scheme) {
			return nil, fmt.Errorf("invalid remote URL: %s", remoteURL)
		}

		return &GitRemote{
			URL:      remoteURL,
			Protocol: gurl.Scheme,
			Host:     gurl.Host,
			Port:     gurl.Port(),
			Path:     gurl.Path,
		}, nil
	}

	if matches := GitRemoteScpRegex.FindStringSubmatch(remoteURL); len(matches) > 0 {
		return &GitRemote{
			URL:      remoteURL,
			Protocol: "ssh",
			User:     matches[GitRemoteScpRegex.SubexpIndex("user")],
			Host:     matches[GitRemoteScpRegex.SubexpIndex("host")],
			Path:     matches[GitRemoteScpRegex.SubexpIndex("path")],
		}, nil
	}

	return nil, fmt.Errorf("invalid remote URL: %s", remoteURL)
}

func ValidateGitRemoteURL(remoteURL string) bool {
	_, err := ParseGitRemoteURL(remoteURL)
	return err == nil
}
