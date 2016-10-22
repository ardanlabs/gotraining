package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func Test_App(t *testing.T) {
	r := require.New(t)
	w := willie.New(App())

	r.Equal(2, len(Customers.All()))

	res := w.Request("/customers").Get()
	body := res.Body.String()
	r.Contains(body, "<h1>Customers</h1>")

	for _, c := range Customers.All() {
		r.Contains(body, fmt.Sprintf("%s - %s", c.ID, c.Name))
	}

	res = w.Request("/customers").Post(url.Values{"Name": []string{"Homer Simpson"}})
	r.Equal(3, len(Customers.All()))
	r.Regexp(`/customers/\d+`, res.Location())

	c, err := Customers.Find("1")
	r.NoError(err)
	res = w.Request("/customers/%s", c.ID).Get()
	body = res.Body.String()
	r.Contains(body, fmt.Sprintf("<h1>%s</h1>", c.Name))
}
