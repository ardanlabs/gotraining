package gothic_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	. "github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/faux"
	"github.com/stretchr/testify/assert"
)

type ProviderStore struct {
	Store map[*http.Request]*sessions.Session
}

func NewProviderStore() *ProviderStore {
	return &ProviderStore{map[*http.Request]*sessions.Session{}}
}

func (p ProviderStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	s := p.Store[r]
	if s == nil {
		s, err := p.New(r, name)
		return s, err
	}
	return s, nil
}

func (p ProviderStore) New(r *http.Request, name string) (*sessions.Session, error) {
	s := sessions.NewSession(p, name)
	s.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 30,
	}
	p.Store[r] = s
	return s, nil
}

func (p ProviderStore) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	p.Store[r] = s
	return nil
}

func init() {
	Store = NewProviderStore()
	goth.UseProviders(&faux.Provider{})
}

func Test_BeginAuthHandler(t *testing.T) {
	a := assert.New(t)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth?provider=faux", nil)
	a.NoError(err)

	BeginAuthHandler(res, req)

	a.Equal(http.StatusTemporaryRedirect, res.Code)
	a.Contains(res.Body.String(),
		`<a href="http://example.com/auth?client_id=&amp;response_type=code&amp;state=state">Temporary Redirect</a>`)
}

func Test_GetAuthURL(t *testing.T) {
	a := assert.New(t)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth?provider=faux", nil)
	a.NoError(err)

	url, err := GetAuthURL(res, req)

	a.NoError(err)

	a.Equal("http://example.com/auth?client_id=&response_type=code&state=state", url)
}

func Test_CompleteUserAuth(t *testing.T) {
	a := assert.New(t)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth/callback?provider=faux", nil)
	a.NoError(err)

	sess := faux.Session{Name: "Homer Simpson", Email: "homer@example.com"}
	session, _ := Store.Get(req, "faux"+SessionName)
	session.Values["faux"] = sess.Marshal()
	err = session.Save(req, res)
	a.NoError(err)

	user, err := CompleteUserAuth(res, req)
	a.NoError(err)

	a.Equal(user.Name, "Homer Simpson")
	a.Equal(user.Email, "homer@example.com")
}

func Test_Logout(t *testing.T) {
	a := assert.New(t)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth/callback?provider=faux", nil)
	a.NoError(err)

	sess := faux.Session{Name: "Homer Simpson", Email: "homer@example.com"}
	session, _ := Store.Get(req, "faux"+SessionName)
	session.Values["faux"] = sess.Marshal()
	err = session.Save(req, res)
	a.NoError(err)

	user, err := CompleteUserAuth(res, req)
	a.NoError(err)

	a.Equal(user.Name, "Homer Simpson")
	a.Equal(user.Email, "homer@example.com")
	err = Logout(res, req)
	a.NoError(err)
	session, _ = Store.Get(req, "faux"+SessionName)
	a.Equal(session.Values, make(map[interface{}]interface{}))
	a.Equal(session.Options.MaxAge, -1)
}

func Test_SetState(t *testing.T) {
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/auth?state=state", nil)
	a.Equal(SetState(req), "state")
}

func Test_GetState(t *testing.T) {
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/auth?state=state", nil)
	a.Equal(GetState(req), "state")
}

func Test_StateValidation(t *testing.T) {
	a := assert.New(t)

	Store = NewProviderStore()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth?provider=faux&state=state_REAL", nil)
	a.NoError(err)

	BeginAuthHandler(res, req)
	session, _ := Store.Get(req, "faux"+SessionName)

	// Assert that matching states will return a nil error
	req, err = http.NewRequest("GET", "/auth/callback?provider=faux&state=state_REAL", nil)
	session.Save(req, res)
	_, err = CompleteUserAuth(res, req)
	a.NoError(err)

	// Assert that mismatched states will return an error
	req, err = http.NewRequest("GET", "/auth/callback?provider=faux&state=state_FAKE", nil)
	session.Save(req, res)
	_, err = CompleteUserAuth(res, req)
	a.Error(err)
}
