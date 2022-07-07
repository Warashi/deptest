package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Warashi/deptest"
	"golang.org/x/tools/cover"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	profile, err := cover.ParseProfiles(filename)
	if err != nil {
		fmt.Println(filepath.Dir(filename))
		return
	}

	packages := deptest.Packages(profile)
	fmt.Println(strings.Join(packages, "\n"))
}
