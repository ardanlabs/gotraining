// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to launch a server with automatic TLS using ACME
// via LetsEncrypt.
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, TLS!")
	})

	log.Println("Listening on port 443")
	log.Println(http.Serve(autocert.NewListener("ardanlabs.com"), nil))
}

// $ curl -v https://ardanlabs.com
// * Connected to ardanlabs.com (45.79.207.220) port 443 (#0)
// * TLS 1.2 connection using TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
// * Server certificate: ardanlabs.com
// * Server certificate: Let's Encrypt Authority X3
// * Server certificate: DST Root CA X3
// > GET / HTTP/1.1
// > Host: ardanlabs.com
// > User-Agent: curl/7.43.0
// > Accept: */*
// >
// < HTTP/1.1 200 OK
// < Date: Mon, 01 May 2017 21:52:28 GMT
// < Content-Length: 12
// < Content-Type: text/plain; charset=utf-8
// <
// Hello, TLS!
