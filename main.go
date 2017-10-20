package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/realityone/git-goget/lib/git"
)

func gopath() []string {
	rawgopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		fmt.Fprintf(os.Stderr, "No $GOPATH environment\n")
		os.Exit(1)
	}
	return filepath.SplitList(rawgopath)
}

func main() {
	gopath := gopath()
	fmt.Printf("%v\n", gopath)

	url, err := git.ParseURL("git@github.com:realityone/git-goget.git")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", url)
}
