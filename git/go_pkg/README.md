# go_pkg

utilities for managing Go packages versions

`go get -u github.com/geektheripper/go-gutils/git/go_pkg`

## Usage

```go
packageMap, err := go_pkg.ResolveVirtualRepo(vrepo)

pkg := go_pkg.Package{
    Name: "git/go_pkg",
    Versions: []*semver.Version{
        semver.MustParse("0.0.1"),
        semver.MustParse("0.0.2"),
        semver.MustParse("0.1.1"),
    },
}

pkg.NextVersion("patch") // 0.1.2
pkg.NextVersion("minor") // 0.2.0
pkg.NextVersion("major") // 1.0.0
```
