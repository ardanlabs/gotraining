package auth

import "testing"

func TestLogin(t *testing.T) {
	_, err := Login("joe", "baz00ka")
	if err != nil {
		t.Fatalf("can't login: %#v", err)
	}
}
