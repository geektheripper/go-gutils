# hack_ssh

get private key that used for specific host and user

implemented by run `ssh -v` and parse the output, so it may not work for all ssh version

`go get -u github.com/geektheripper/go-gutils/git/hack_ssh`

## Usage

```go
key, err := hack_ssh.ResolveHost("github.com", "git")
```
