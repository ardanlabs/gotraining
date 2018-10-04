package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client knows how to call the GitHub API to get contributor information.
type Client struct {
	token  string
	client http.Client
}

// NewClient builds a Client value. It validates that the token field is set
// and is not blank. It sets internal state for the http client. Call it like:
//	github.NewClient(os.Getenv("GITHUB_TOKEN"))
func NewClient(token string) (*Client, error) {

	if token == "" {
		return nil, errors.New("token is required")
	}

	return &Client{
		token:  token,
		client: http.Client{Timeout: 5 * time.Second},
	}, nil
}

// ContributorList gives a list of the top 30 contributors. It returns an error
// for network problems reaching the API or for application problems such as a
// 404 or 403 response from GitHub.
func (c *Client) ContributorList(repo string) ([]Contributor, error) {

	// Make a request and set the auth token in the header.
	url := fmt.Sprintf("https://api.github.com/repos/%s/contributors", repo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	// Ensure response has a 200 status code.
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
