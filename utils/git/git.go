package git

import (
	"net/url"
	"regexp"
	"strings"
)

// URL contains all parsed information
type URL struct {
	Protocol string
	Host     string
	Owner    string
	RepoPath string
}

var hasSchemePattern = regexp.MustCompile("^[^:]+://")
var scpLikeURLPattern = regexp.MustCompile("^([^@]+)@?([^:]+):/?(.+)$")

func trimPrefix(path string) string {
	if strings.HasPrefix(path, "/") {
		return strings.TrimLeft(path, "/")
	}
	return path
}

// ParseURL will parse giving url into `git.URL` structure
func ParseURL(rawurl string) (*URL, error) {
	if !hasSchemePattern.MatchString(rawurl) && scpLikeURLPattern.MatchString(rawurl) {
		matched := scpLikeURLPattern.FindStringSubmatch(rawurl)
		user := matched[1]
		host := matched[2]
		path := matched[3]

		return &URL{
			Protocol: "ssh",
			Host:     host,
			Owner:    user,
			RepoPath: trimPrefix(path),
		}, nil
	}

	parsed, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	owner := ""
	if parsed.User != nil {
		owner = parsed.User.Username()
	}

	return &URL{
		Protocol: parsed.Scheme,
		Host:     parsed.Host,
		Owner:    owner,
		RepoPath: trimPrefix(parsed.EscapedPath()),
	}, nil
}

// SplitRepoPath will split repo path into a list of directories
func (u *URL) SplitRepoPath() []string {
	return strings.Split(u.RepoPath, "/")
}
