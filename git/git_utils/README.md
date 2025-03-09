# git_utils

utils for git, used go standard library only

`go get -u github.com/geektheripper/go-gutils/git/git_utils`

## Usage

```go
remote, err := git_utils.ParseGitRemoteURL("https://github.com/user/repo.git")
ok := git_utils.ValidateGitRemoteURL("redis://localhost:6379/0")
```
