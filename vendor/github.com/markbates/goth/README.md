# Goth: Multi-Provider Authentication for Go [![GoDoc](https://godoc.org/github.com/markbates/goth?status.svg)](https://godoc.org/github.com/markbates/goth) [![Build Status](https://travis-ci.org/markbates/goth.svg)](https://travis-ci.org/markbates/goth)

Package goth provides a simple, clean, and idiomatic way to write authentication
packages for Go web applications.

Unlike other similar packages, Goth, lets you write OAuth, OAuth2, or any other
protocol providers, as long as they implement the `Provider` and `Session` interfaces.

This package was inspired by [https://github.com/intridea/omniauth](https://github.com/intridea/omniauth).

## Installation

```text
$ go get github.com/markbates/goth
```

## Supported Providers

* Amazon
* Auth0
* Battle.net
* Bitbucket
* Box
* Cloud Foundry
* Dailymotion
* Deezer
* Digital Ocean
* Discord
* Dropbox
* Facebook
* Fitbit
* GitHub
* Gitlab
* Google+
* Heroku
* InfluxCloud
* Instagram
* Intercom
* Lastfm
* Linkedin
* Meetup
* OneDrive
* OpenID Connect (auto discovery)
* Paypal
* SalesForce
* Slack
* Soundcloud
* Spotify
* Steam
* Stripe
* Twitch
* Twitter
* Uber
* Wepay
* Xero
* Yahoo
* Yammer

## Examples

See the [examples](examples) folder for a working application that lets users authenticate
through Twitter, Facebook, Google Plus etc.

To run the example either clone the source from GitHub

```text
$ git clone git@github.com:markbates/goth.git
```
or use
```text
$ go get github.com/markbates/goth
```
```text
$ cd goth/examples
$ go get -v
$ go build 
$ ./examples
```

Now open up your browser and go to [http://localhost:3000](http://localhost:3000) to see the example.

To actually use the different providers, please make sure you set environment variables. Example given in the examples/main.go file

## Security Notes

By default, gothic uses a `CookieStore` from the `gorilla/sessions` package to store session data.

As configured, this default store (`gothic.Store`) will generate cookies with `Options`:

```go
&Options{
   Path:   "/",
   Domain: "",
   MaxAge: 86400 * 30,
   HttpOnly: true,
   Secure: false,
 }
```

To tailor these fields for your application, you can override the `gothic.Store` variable at startup.

The follow snippet show one way to do this:

```go
key := ""             // Replace with your SESSION_SECRET or similar
maxAge := 86400 * 30  // 30 days
isProd := false       // Set to true when serving over https

store := sessions.NewCookieStore([]byte(key))
store.MaxAge(maxAge)
store.Options.Path = "/"
store.Options.HttpOnly = true   // HttpOnly should always be enabled
store.Options.Secure = isProd

gothic.Store = store
```

## Issues

Issues always stand a significantly better chance of getting fixed if they are accompanied by a
pull request.

## Contributing

Would I love to see more providers? Certainly! Would you love to contribute one? Hopefully, yes!

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## Contributors

* Mark Bates
* Tyler Bunnell
* Corey McGrillis
* willemvd
* Rakesh Goyal
* Andy Grunwald
* Glenn Walker
* Kevin Fitzpatrick
* Ben Tranter
* Sharad Ganapathy
* Andrew Chilton
* sharadgana
* Aurorae
* Craig P Jolicoeur
* Zac Bergquist
* Geoff Franks
* Raphael Geronimi
* Noah Shibley
* lumost
* oov
* Felix Lamouroux
* Rafael Quintela
* Tyler
* DenSm
* Samy KACIMI
* dante gray
* Noah
* Jacob Walker
* Marin Martinic
* Roy
* Omni Adams
* Sasa Brankovic
* dkhamsing
* Dante Swift
* Attila Domokos
* Albin Gilles
* Syed Zubairuddin
* Johnny Boursiquot
* Jerome Touffe-Blin
* bryanl
* Masanobu YOSHIOKA
* Jonathan Hall
* HaiMing.Yin
* Sairam Kunala
* Regan Ashworth
