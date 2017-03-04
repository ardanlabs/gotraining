## Web - TLS

Learn about securing your application using TLS.

## Notes

* Usually apps just listen for HTTP and offload TLS termination to a load balancer like Caddy or Nginx.
* The net/http package provides support for TLS if you really need it
* The crypto/tls package comes with a program for generating self signed certificates

    go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost

## Links

https://golang.org/pkg/net/http/  
https://golang.org/pkg/crypto/tls/  
https://caddyserver.com/  

## Code Review

TLS support: [Code](example1/main.go)

## Exercises

### Exercise 1

TBD
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
