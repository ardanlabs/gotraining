package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	jose "github.com/dvsekhvalnov/jose2go"
)

type JWTTransporter struct {
	transporter  http.RoundTripper
	sharedSecret []byte
}

func (c *JWTTransporter) RoundTrip(req *http.Request) (*http.Response, error) {
	buf, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	token, err := jose.Sign(string(buf), jose.HS256, c.sharedSecret)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-signature", token)
	res, err := c.transporter.RoundTrip(req)
	return res, err
}
