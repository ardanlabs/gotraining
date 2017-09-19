package github_test

import (
	"testing"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/stretchr/testify/assert"
)

func Test_Implements_Session(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &github.Session{}

	a.Implements((*goth.Session)(nil), s)
}

func Test_GetAuthURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &github.Session{}

	_, err := s.GetAuthURL()
	a.Error(err)

	s.AuthURL = "/foo"

	url, _ := s.GetAuthURL()
	a.Equal(url, "/foo")
}

func Test_ToJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &github.Session{}

	data := s.Marshal()
	a.Equal(data, `{"AuthURL":"","AccessToken":""}`)
}

func Test_String(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &github.Session{}

	a.Equal(s.String(), s.Marshal())
}
