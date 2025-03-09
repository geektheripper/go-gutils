# virtual_repo

based `go-git/go-git`

create an in memory virtual repo, and provide functions to build tag and push

`go get -u github.com/geektheripper/go-gutils/git/virtual_repo`

## Usage

```go
repo, err := virtual_repo.NewVirtualRepo("https://github.com/geektheripper/go-gutils.git", "working-branch")
if err != nil {
	panic(err)
}

repo.CreateFile("README.md", strings.NewReader("Hello, World!"), 0644)
repo.PushTag("v0.0.1", "Initial commit")

repo.DeleteRemoteTag("v0.0.1")
```
