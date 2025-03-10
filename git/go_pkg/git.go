package go_pkg

import (
	"regexp"

	"github.com/Masterminds/semver"
	"github.com/geektheripper/go-gutils/git/virtual_repo"
)

var packageRefRegex = regexp.MustCompile(`^refs/tags/(?P<package>.*)/v(?P<version>.*)$`)

func ResolveVirtualRepo(v *virtual_repo.VirtualRepo) (map[string]*Package, error) {
	refs, err := v.FilterRefs("refs/tags/")
	if err != nil {
		return nil, err
	}

	packages := map[string]*Package{}
	for _, ref := range refs {
		tag := ref.Name().String()
		matches := packageRefRegex.FindStringSubmatch(tag)

		if len(matches) == 0 {
			continue
		}

		packageName := matches[packageRefRegex.SubexpIndex("package")]
		versionText := matches[packageRefRegex.SubexpIndex("version")]

		if _, ok := packages[packageName]; !ok {
			packages[packageName] = &Package{
				Name:     packageName,
				Versions: []*semver.Version{},
			}
		}

		if version, err := semver.NewVersion(versionText); err != nil {
			continue
		} else {
			packages[packageName].Versions = append(packages[packageName].Versions, version)
		}
	}

	return packages, nil
}
