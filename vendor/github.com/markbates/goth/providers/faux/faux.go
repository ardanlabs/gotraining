// Package faux is used exclusively for testing purposes. I would strongly suggest you move along
// as there's nothing to see here.
package faux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/markbates/goth"
	"golang.org/x/oauth2"
)

// Provider is used only for testing.
type Provider struct {
	HTTPClient   *http.Client
	providerName string
}

// Session is used only for testing.
type Session struct {
	ID          string
	Name        string
	Email       string
	AuthURL     string
	AccessToken string
}

// Name is used only for testing.
func (p *Provider) Name() string {
	return "faux"
}

// SetName is to update the name of the provider (needed in case of multiple providers of 1 type)
func (p *Provider) SetName(name string) {
	p.providerName = name
}

// BeginAuth is used only for testing.
func (p *Provider) BeginAuth(state string) (goth.Session, error) {
	c := &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL: "http://example.com/auth",
		},
	}
	url := c.AuthCodeURL(state)
	return &Session{
		ID:      "id",
		AuthURL: url,
	}, nil
}

// FetchUser is used only for testing.
func (p *Provider) FetchUser(session goth.Session) (goth.User, error) {
	sess := session.(*Session)
	user := goth.User{
		UserID:      sess.ID,
		Name:        sess.Name,
		Email:       sess.Email,
		Provider:    p.Name(),
		AccessToken: sess.AccessToken,
	}

	if user.AccessToken == "" {
		return user, fmt.Errorf("%s cannot get user information without accessToken", p.providerName)
	}
	return user, nil
}

// UnmarshalSession is used only for testing.
func (p *Provider) UnmarshalSession(data string) (goth.Session, error) {
	sess := &Session{}
	err := json.NewDecoder(strings.NewReader(data)).Decode(sess)
	return sess, err
}

func (p *Provider) Client() *http.Client {
	return goth.HTTPClientWithFallBack(p.HTTPClient)
}

// Debug is used only for testing.
func (p *Provider) Debug(debug bool) {}

//RefreshTokenAvailable is used only for testing
func (p *Provider) RefreshTokenAvailable() bool {
	return false
}

//RefreshToken is used only for testing
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	return nil, nil
}

// Authorize is used only for testing.
func (s *Session) Authorize(provider goth.Provider, params goth.Params) (string, error) {
	s.AccessToken = "access"
	return s.AccessToken, nil
}

// Marshal is used only for testing.
func (s *Session) Marshal() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetAuthURL is used only for testing.
func (s *Session) GetAuthURL() (string, error) {
	return s.AuthURL, nil
}
