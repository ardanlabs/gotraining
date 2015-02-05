package tests

import (
	"testing"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
)

const succeed = "\xE2\x9C\x93"
const failed = "\xE2\x9C\x97"

var c *app.Context

// TestMain is the entry point for the test. Used to create a context
// before the tests are run and can then perform cleanup.
func TestMain(m *testing.M) {
	c = &app.Context{
		Session:   app.GetSession(),
		SessionID: "TESTING",
	}

	m.Run()

	c.Session.Close()
}
