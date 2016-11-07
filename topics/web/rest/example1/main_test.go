// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// NOTE: This feels wrong to me?

// Tests for the sample program to show how to use a regex
// to handle REST based URL schemas and routes.
package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/ardanlabs/gotraining/topics/web/customer"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {

	// Create a new assertion.
	r := require.New(t)

	// Connect the routes into willie.
	w := willie.New(App())

	// Do we have two customers.
	r.Equal(2, len(customer.All()))

	// Perform a GET request for customers.
	res := w.Request("/customers").Get()

	// Does the response contain this string?
	body := res.Body.String()
	r.Contains(body, "<h1>Customers</h1>")

	// Check that all the customers are represented
	// in the response.
	for _, c := range customer.All() {
		r.Contains(body, fmt.Sprintf("%d - %s", c.ID, c.Name))
	}

	// Perform a POST request to save a customer.
	res = w.Request("/customers").Post(url.Values{"Name": []string{"Homer Simpson"}})

	// Check that we now how three customers.
	r.Equal(3, len(customer.All()))

	// Check we got the correct redirect url.
	r.Regexp(`/customers/\d+`, res.Location())

	// Find custoner 1 in the DB.
	c, err := customer.Find(1)
	r.NoError(err)

	// Perform a GET request for customer 1.
	res = w.Request("/customers/%d", c.ID).Get()

	// Does the response contain this string?
	body = res.Body.String()
	r.Contains(body, fmt.Sprintf("<h1>%s</h1>", c.Name))
}
