package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Warashi/deptest"
	"golang.org/x/tools/cover"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	profile, err := cover.ParseProfiles(filename)
	if err != nil {
		log.Fatalf("cover.ParseProfiles(%q): %v", filename, err)
	}

	packages := deptest.Packages(profile)
	fmt.Println(strings.Join(packages, "\n"))
}
