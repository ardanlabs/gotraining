// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.3

// We only compile this with Go 1.3+ because previously tls.DialWithDialer
// wasn't available. The dial.go file is used for previous Go versions.

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
		conn, err = tls.DialWithDialer(dialer, "tcp", parseAuthority(config.Location), config.TlsConfig)

	default:
		err = ErrBadScheme
	}
	return
}
