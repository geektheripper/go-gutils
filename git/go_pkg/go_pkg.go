package go_pkg

import (
	"sort"

	"github.com/Masterminds/semver"
)

type Package struct {
	Name     string
	Versions []*semver.Version
}

func (p *Package) GetLatestVersion() *semver.Version {
	sort.Sort(semver.Collection(p.Versions))
	return p.Versions[len(p.Versions)-1]
}

func (p *Package) NextVersion(upgradeType ...string) *semver.Version {
	if len(upgradeType) > 1 {
		panic("only one upgrade type is allowed")
	}

	latestVersion := p.GetLatestVersion()

	nextVersion := latestVersion.IncPatch()

	switch upgradeType[0] {
	case "patch":
		nextVersion = latestVersion.IncPatch()
	case "minor":
		nextVersion = latestVersion.IncMinor()
	case "major":
		nextVersion = latestVersion.IncMajor()
	}

	return &nextVersion
}
