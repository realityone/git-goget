package git_test

import (
	"testing"

	"github.com/realityone/git-goget/lib/git"
)

func TestParseGitURL(t *testing.T) {
	testUrls := map[string]*git.URL{
		"git@github.com:realityone/git-goget.git": &git.URL{
			Protocol: "ssh",
			Host:     "github.com",
			Owner:    "git",
			RepoPath: "realityone/git-goget.git",
		},
		"https://github.com/realityone/git-goget.git": &git.URL{
			Protocol: "https",
			Host:     "github.com",
			Owner:    "",
			RepoPath: "realityone/git-goget.git",
		},
	}
	for rawurl, giturl := range testUrls {
		parsedurl, err := git.ParseURL(rawurl)
		if err != nil {
			t.Error(err)
			continue
		}
		if *parsedurl != *giturl {
			t.Errorf("parsed url %v is not equal to defined %v\n", parsedurl, giturl)
			continue
		}
	}
}
