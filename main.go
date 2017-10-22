package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/realityone/git-goget/utils/git"
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

func fetchRepo(rawurl, gopath string, requireUpdate bool) {
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

	clone := func() {
		cmd := exec.Command("sh", "-c", fmt.Sprintf(`
			git clone %s
			`, path))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to clone repository `%s`: %s\n", rawurl, err)
			os.Exit(1)
		}
	}

	update := func() {
		cmd := exec.Command("sh", "-c", fmt.Sprintf(`
			cd %s && \
			git pull && \
			git reset --hard HEAD
			`, path))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update repository `%s`: %s\n", rawurl, err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(filepath.Join(path, ".git")); !os.IsNotExist(err) && requireUpdate {
		update()
		return
	}

	clone()
}

func usage() {
	fmt.Println("usage: git goget [-u] [REPOSITORY...]")
	os.Exit(1)
}

func main() {
	gopath := gopath()

	update := flag.Bool("u", false, "update the package")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	for _, repoURL := range args {
		fetchRepo(repoURL, gopath[0], *update)
	}
}
