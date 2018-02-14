package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// Contributor summarizes one person's contributions to a particular
// GitHub repository.
type Contributor struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

// Client knows how to call the GitHub API to get contributor information.
type Client struct {
	token  string
	client http.Client

	baseURL string // The root URL of the API. We make this an unexported field so we can override it in tests.
}

// tokenRE defines the form of a valid token. We compile it once at package
// load time with MustCompile because we know it will never change and the
// regexp is always valid.
var tokenRE = regexp.MustCompile(`^[0-9a-f]{40}$`)

// NewClient builds a Client value. It validates that the token field is set
// and is of a valid form. It sets internal state for the http client.
// Call it like:
//	github.NewClient(os.Getenv("GITHUB_TOKEN"))
func NewClient(token string) (*Client, error) {

	if token == "" {
		return nil, errors.New("token is required")
	}
	if !tokenRE.MatchString(token) {
		return nil, errors.New("token is required")
	}

	return &Client{
		token:   token,
		client:  http.Client{Timeout: 5 * time.Second},
		baseURL: "https://api.github.com", // TODO pass this in
	}, nil
}

// repoRE is the regexp value for checking repo strings. We compile this once
// with "MustCompile" when the package loads because it will never change and
// we know it will always work.
var repoRE = regexp.MustCompile(`[a-zA-Z0-9]+/[a-zA-Z0-9]+`)

// Contributors gives a list of the top 30 contributors. It returns an error
// for network problems reaching the API or for application problems such as a
// 404 or 403 response from GitHub.
func (c *Client) Contributors(repo string) ([]Contributor, error) {
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
