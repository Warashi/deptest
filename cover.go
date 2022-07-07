package deptest

import (
	"path/filepath"

	"golang.org/x/exp/slices"
	"golang.org/x/tools/cover"
)

func Packages(profiles []*cover.Profile) []string {
	packages := make([]string, 0, len(profiles))
	for _, p := range profiles {
		packages = append(packages, filepath.Dir(p.FileName))
	}
	slices.Sort(packages)
	return slices.Compact(packages)
}
