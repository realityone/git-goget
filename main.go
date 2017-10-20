package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func goRepoPath(gopath string, url *git.URL) string {
	repoPath := url.RepoPath
	if strings.HasSuffix(url.RepoPath, ".git") {
		repoPath = strings.TrimSuffix(url.RepoPath, ".git")
	}
	return filepath.Join(gopath, "src", url.Host, repoPath)
}

func cloneRepo(rawurl, gopath string) {
	url, err := git.ParseURL(rawurl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse url `%s`: %s\n", rawurl, err)
		os.Exit(1)
	}

	path := goRepoPath(gopath, url)
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory on `%s`: %s\n", path, err)
		os.Exit(1)
	}

	cloneCmd := exec.Command("git", "clone", rawurl, path)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clone repository `%s`: %s\n", rawurl, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("git goget [REPOSITORY...]")
	os.Exit(1)
}

func main() {
	gopath := gopath()
	if len(os.Args) <= 1 {
		usage()
	}

	for _, repoURL := range os.Args[1:] {
		cloneRepo(repoURL, gopath[0])
	}
}
