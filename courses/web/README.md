## Ultimate Web
Ultimate Web is a 2 day class for any Go developer who wishes to learn how to build robust and well tested HTTP based applications in Go. This class provides an intensive, comprehensive and idiomatic view build Web, SOA, and API applications using Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## HTTP Basics
This section isn’t really about “Go”. It should cover the basics of HTTP. What is a request, a response, status codes, basic theory behind HTTP.

* What is HTTP?
* How does it work?

## HTTP Basics in Go
With a basic understand of how the web and HTTP work, let’s write a simple “Hello World” app in Go. We’ll cover how to start a web server in Go, take in requests, and return responses.

* Hello World
* http.ServeMux
* http.ListenAndServe
* http.Handler
* http.HandlerFunc

## Testing HTTP in Go
Now that we have some code written, let’s start understanding how to test HTTP Go apps. We’ll look at two different ways of testing HTTP apps.

* Direct testing of http.Handlers (Unit)
* Testing with httptest.NewServer (Integration)

## HTML Templates
We can now write primitive web apps, as well as test them. Now we need to start adding some meat to it. This section covers generating HTML using Go templates, as well as how to serve up static files, and how to bundle those files in a finished binary.

* Go Templates
* Assets (Static Files)
* Project/File Structure
* Bundling Assets in Binaries (go.rice)

## POST Requests
HTTP applications don’t just serve content, they also take in content. We’ll branch out of GET requests and start taking in POST requests, processing forms, handling file uploads, and of course, how to test all of this.

* Form Handling (Handling POST requests)
* File Uploads

## Sessions and Cookies
Managing sessions and cookies is an important part of every web application. Whether it's keeping a user "logged in" or tracking who visits your site, these concepts are essential to learn.

* Sessions
* Cookies

## Introduction to REST
The app is starting to get more complex, at this point we should start talking about design patterns around building web applications, in particular we’ll discuss RESTful design.

* URL Design
* More HTTP Verbs
* Error Handling

## Alternative Muxers
The basic muxer in Go has gotten us a long way by this point, but it has it’s limitations. Let’s tour three very different types muxers/routers.

* Gorilla (compatible with the standard library)
* Echo (a micro-framework)
* Httprouter (non-compatible with the standard library)

## APIs
By this point we should be able to build fully featured HTML applications in Go, but the fun doesn’t stop there. Let’s turn an eye to building APIs for other applications to consume. We’ll look at two different ways of building APIs (RESTful & HyperMedia), we’ll also look at different ways to handle that age old question of versioning and API.

* RESTful
* HyperMedia
* Versioning

## Data Serialization
An API is nothing if we don’t have a way of serializing the data we want to send. We’ll look at 3 common data formats, as well as ways to customize those formats to match the needs of your API.

* JSON
* XML
* ProtoBufs

## Consuming HTTP APIs
What good is having an API if we can’t consume it? We’ll learn how to use Go to speak with APIs, marshal & unmarshal data, set request headers, and more.

* http.Client

## Securing APIs
Secure APIs are a must have these days. From configuring our server to use SSL to learning how to read/write JSON Web Tokens we’ll learn how to trust communication between web servers.

* JWT
* SSL

## Context
In Go 1.7 there is now a Context value that can passed along with requests that make working with multiple services, and go routines, that much more pleasant. Let’s looking into this mystery new addition and see how we can use it to set/get values through a request, as well as control requests to 3rd parties as well.

* Introduction
* Basic Usage
* Using w/ Multiple Services

## Deployment
We’ve built our application and we’re ready to release it to the world! But how? We’ll cover a simple deployment strategy using Heroku. We’ll also look at the Caddy Web Server (written in Go) that has been taking the world by storm, and see just how easy it is to get started with Caddy and Go web applications.

* Heroku
* Caddy Web Server
