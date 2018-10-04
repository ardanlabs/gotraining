package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// API is the root URL for the github.com api. Use this with NewClient.
const API = "https://api.github.com"

// Client knows how to call the GitHub API to get contributor information.
type Client struct {
	token   string
	client  http.Client
	baseURL string
}

// tokenRE defines the form of a valid token. We compile it once at package
// load time with MustCompile because we know it will never change and the
// regexp is always valid.
var tokenRE = regexp.MustCompile(`^[0-9a-f]{40}$`)

// NewClient builds a Client value. It needs the root url for the API. Use the
// const API for github.com or pass your own url for tests or an enterprise
// installation. NewClient will error if the token field is invalid.
// Call it like:
//	github.NewClient(github.API, os.Getenv("GITHUB_TOKEN"))
func NewClient(root, token string) (*Client, error) {

	if token == "" {
		return nil, errors.New("token is required")
	}
	if !tokenRE.MatchString(token) {
		return nil, errors.New("token is invalid")
	}

	return &Client{
		token:   token,
		client:  http.Client{Timeout: 5 * time.Second},
		baseURL: root,
	}, nil
}

// repoRE is the regexp value for checking repo strings. We compile this once
// with "MustCompile" when the package loads because it will never change and
// we know it will always work.
var repoRE = regexp.MustCompile(`[a-zA-Z0-9]+/[a-zA-Z0-9]+`)

// ContributorList gives a list of the top 30 contributors. It returns an error
// for network problems reaching the API or for application problems such as a
// 404 or 403 response from GitHub.
func (c *Client) ContributorList(repo string) ([]Contributor, error) {
	if repo == "" {
		return nil, errors.New("repo is required")
	}
	if !repoRE.MatchString(repo) {
		return nil, errors.New("repo is invalid")
	}

	// Make a request and set the auth token in the header.
	u := fmt.Sprintf("%s/repos/%s/contributors", c.baseURL, repo)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	// Execute the request.
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API responded with a %d %s", resp.StatusCode, resp.Status)
	}

	// Decode the result.
	var cons []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&cons); err != nil {
		return nil, err
	}

	return cons, nil
}
