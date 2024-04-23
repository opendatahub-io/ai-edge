package support

import (
	"fmt"
	"net/url"
	"strings"
)

type GitRepoURL struct {
	Server   string
	OrgName  string
	RepoName string
}

func ParseGitURL(rawURL string) (GitRepoURL, error) {
	gitRepoURL := GitRepoURL{}

	URL, err := url.Parse(rawURL)
	if err != nil {
		return gitRepoURL, err
	}

	// using trim here aswell because leading and trailing / causes empty strings when slicing
	subPaths := strings.Split(strings.Trim(URL.Path, "/"), "/")
	if len(subPaths) != 2 {
		return gitRepoURL, fmt.Errorf("cannot parse git repo URL, expected [scheme]://[host]/[org]/[repo] got %v", rawURL)
	}

	gitRepoURL.OrgName = subPaths[0]
	gitRepoURL.RepoName = subPaths[1]
	gitRepoURL.Server = fmt.Sprintf("%v://%v", URL.Scheme, URL.Host)

	return gitRepoURL, nil
}
