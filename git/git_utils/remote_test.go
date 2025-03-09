package git_utils_test

import (
	"strings"
	"testing"

	"github.com/geektheripper/go-gutils/git/git_utils"
)

type GitRemoteTestCase struct {
	name        string
	remoteURL   string
	wantRemote  *git_utils.GitRemote
	wantErr     bool
	errContains string
}

func TestParseGitRemoteURL(t *testing.T) {
	tests := []GitRemoteTestCase{
		{
			name:      "valid https URL",
			remoteURL: "https://github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "https://github.com/user/repo.git",
				Protocol: "https",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid http URL",
			remoteURL: "http://github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "http://github.com/user/repo.git",
				Protocol: "http",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid ssh URL",
			remoteURL: "ssh://git@github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "ssh://git@github.com/user/repo.git",
				Protocol: "ssh",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid git URL",
			remoteURL: "git://github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "git://github.com/user/repo.git",
				Protocol: "git",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid ftp URL",
			remoteURL: "ftp://github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "ftp://github.com/user/repo.git",
				Protocol: "ftp",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid ftps URL",
			remoteURL: "ftps://github.com/user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "ftps://github.com/user/repo.git",
				Protocol: "ftps",
				Host:     "github.com",
				Path:     "/user/repo.git",
			},
			wantErr: false,
		},
		{
			name:      "valid SCP format",
			remoteURL: "git@github.com:user/repo.git",
			wantRemote: &git_utils.GitRemote{
				URL:      "git@github.com:user/repo.git",
				Protocol: "ssh",
				User:     "git",
				Host:     "github.com",
				Path:     "user/repo.git",
			},
			wantErr: false,
		},
		{
			name:        "URL without .git suffix",
			remoteURL:   "https://github.com/user/repo",
			wantRemote:  nil,
			wantErr:     true,
			errContains: "must end with .git",
		},
		{
			name:        "invalid protocol",
			remoteURL:   "invalid://github.com/user/repo.git",
			wantRemote:  nil,
			wantErr:     true,
			errContains: "invalid remote URL",
		},
		{
			name:        "invalid SCP format",
			remoteURL:   "invalid@format:user/repo.git",
			wantRemote:  nil,
			wantErr:     true,
			errContains: "invalid remote URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := git_utils.ParseGitRemoteURL(tt.remoteURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGitRemoteURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errContains != "" && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ParseGitRemoteURL() error = %v, should contain %v", err, tt.errContains)
				}
				return
			}
			if got == nil && tt.wantRemote == nil {
				return
			}
			if (got == nil) != (tt.wantRemote == nil) {
				t.Errorf("ParseGitRemoteURL() got = %v, want %v", got, tt.wantRemote)
				return
			}
			if got.URL != tt.wantRemote.URL {
				t.Errorf("ParseGitRemoteURL() URL = %v, want %v", got.URL, tt.wantRemote.URL)
			}
			if got.Protocol != tt.wantRemote.Protocol {
				t.Errorf("ParseGitRemoteURL() Protocol = %v, want %v", got.Protocol, tt.wantRemote.Protocol)
			}
			if got.User != tt.wantRemote.User {
				t.Errorf("ParseGitRemoteURL() User = %v, want %v", got.User, tt.wantRemote.User)
			}
			if got.Host != tt.wantRemote.Host {
				t.Errorf("ParseGitRemoteURL() Host = %v, want %v", got.Host, tt.wantRemote.Host)
			}
			if got.Path != tt.wantRemote.Path {
				t.Errorf("ParseGitRemoteURL() Path = %v, want %v", got.Path, tt.wantRemote.Path)
			}
		})
	}
}

type ValidateGitRemoteURLTestCase struct {
	name      string
	remoteURL string
	want      bool
}

func TestValidateGitRemoteURL(t *testing.T) {
	tests := []ValidateGitRemoteURLTestCase{
		{
			name:      "valid URL",
			remoteURL: "https://github.com/user/repo.git",
			want:      true,
		},
		{
			name:      "valid SCP format",
			remoteURL: "git@github.com:user/repo.git",
			want:      true,
		},
		{
			name:      "invalid URL",
			remoteURL: "https://github.com/user/repo",
			want:      false,
		},
		{
			name:      "invalid protocol",
			remoteURL: "invalid://github.com/user/repo.git",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := git_utils.ValidateGitRemoteURL(tt.remoteURL); got != tt.want {
				t.Errorf("ValidateGitRemoteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
