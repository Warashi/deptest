package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"golang.org/x/tools/cover"

	"github.com/Warashi/deptest"
)

var module string

//go:embed makefile.tmpl
var tmpl string

func init() {
	flag.StringVar(&module, "module", "", "module name")
}

type Template struct {
	Profile string
	Deps    []string
}

func main() {
	flag.Parse()

	t := make([]Template, 0, flag.NArg())
	for _, arg := range flag.Args() {
		profile, err := cover.ParseProfiles(arg)
		var files []string
		files, err = filepath.Glob(filepath.Dir(arg) + "/*.go")
		if err != nil {
			log.Fatalln(err)
		}

		deps := deptest.Files(profile)
		for i := range deps {
			deps[i] = strings.TrimPrefix(deps[i], module+"/")
		}

		files = slices.Grow(files, len(deps))
		for _, dep := range deps {
			if _, err := os.Stat(dep); err == nil {
				files = append(files, dep)
			}
		}

		slices.Sort(files)
		files = slices.Compact(files)

		t = append(t, Template{Profile: arg, Deps: files})
	}

	tmpl, err := template.New("makefile").Parse(tmpl)
	if err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.Execute(os.Stdout, t); err != nil {
		log.Fatalln(err)
	}
}
