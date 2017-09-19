package github_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := githubProvider()
	a.Equal(provider.ClientKey, os.Getenv("GITHUB_KEY"))
	a.Equal(provider.Secret, os.Getenv("GITHUB_SECRET"))
	a.Equal(provider.CallbackURL, "/foo")
}

func Test_NewCustomisedURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := urlCustomisedURLProvider()
	session, err := p.BeginAuth("test_state")
	s := session.(*github.Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "http://authURL")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Implements((*goth.Provider)(nil), githubProvider())
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := githubProvider()
	session, err := provider.BeginAuth("test_state")
	s := session.(*github.Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "github.com/login/oauth/authorize")
	a.Contains(s.AuthURL, fmt.Sprintf("client_id=%s", os.Getenv("GITHUB_KEY")))
	a.Contains(s.AuthURL, "state=test_state")
	a.Contains(s.AuthURL, "scope=user")
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := githubProvider()

	s, err := provider.UnmarshalSession(`{"AuthURL":"http://github.com/auth_url","AccessToken":"1234567890"}`)
	a.NoError(err)
	session := s.(*github.Session)
	a.Equal(session.AuthURL, "http://github.com/auth_url")
	a.Equal(session.AccessToken, "1234567890")
}

func githubProvider() *github.Provider {
	return github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "/foo", "user")
}

func urlCustomisedURLProvider() *github.Provider {
	return github.NewCustomisedURL(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "/foo", "http://authURL", "http://tokenURL", "http://profileURL", "http://emailURL")
}
