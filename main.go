package main

import (
	"fmt"
	"os"
	"path/filepath"
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
}
