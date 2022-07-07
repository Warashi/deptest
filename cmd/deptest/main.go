package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
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
		if errors.Is(os.ErrNotExist, err) {
			fmt.Println(filepath.Dir(filename))
			return
		}
		log.Fatalf("cover.ParseProfiles(%q): %v", filename, err)
	}

	packages := deptest.Packages(profile)
	fmt.Println(strings.Join(packages, "\n"))
}
