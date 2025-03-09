# hack_ssh

find out which private key can login to a host, implemented by actually run `ssh -v` and parse the output, so it may not work for all ssh version

`go get -u github.com/geektheripper/go-gutils/git/hack_ssh`

## Usage

```go
key, err := hack_ssh.GetKeyForHost("github.com", "git")
```
