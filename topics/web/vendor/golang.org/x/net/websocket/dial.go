// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.3

package websocket

import (
	"crypto/tls"
	"net"
)

func dialWithDialer(dialer *net.Dialer, config *Config) (conn net.Conn, err error) {
	switch config.Location.Scheme {
	case "ws":
		conn, err = dialer.Dial("tcp", parseAuthority(config.Location))

	case "wss":
		conn, err = tls.Dial("tcp", parseAuthority(config.Location), config.TlsConfig)

	default:
		err = ErrBadScheme
	}
	return
}
