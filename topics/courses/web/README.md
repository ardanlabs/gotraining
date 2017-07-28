## Ultimate Web

This is material for any Go developer who wishes to learn how to build robust and well tested HTTP based applications in Go. This class provides an intensive, comprehensive and idiomatic view of how to build Web, SOA, and API applications using Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

[Design Document](../../web/README.md)

## HTTP Basics in Go

With a basic understanding of how the web and HTTP work, let’s write a simple “Hello World” app in Go. We’ll cover how to start a web server in Go, take in requests, and return responses.

[Basics](../../web/basics/README.md)

## Testing HTTP in Go

Now that we have some code written, let’s start understanding how to test HTTP Go apps. We’ll look at two different ways of testing HTTP apps.

[Testing](../../web/testing/README.md)

## POST Requests

HTTP applications don’t just serve content, they also take in content. We’ll branch out of GET requests and start taking in POST requests, processing forms, handling file uploads, and of course, how to test all of this.

[Post Requests](../../web/posts/README.md)

## HTML Templates

We can now write primitive web apps, as well as test them. Now we need to start adding some meat to it. This section covers generating HTML using Go templates, as well as how to serve up static files, and how to bundle those files in a finished binary.

[Templates](../../web/templates/README.md)

## Sessions and Cookies

Managing sessions and cookies is an important part of every web application. Whether it's keeping a user "logged in" or tracking who visits your site, these concepts are essential to learn.

[Sessions and Cookies](../../web/sessions_cookies/README.md)

## Introduction to REST

The app is starting to get more complex, at this point we should start talking about design patterns around building web applications, in particular we’ll discuss RESTful design.

[REST](../../web/rest/README.md)

## Alternative Muxers

The basic muxer in Go has gotten us a long way by this point, but it has its limitations. Let’s tour three very different types of muxers/routers.

[Muxing](../../web/muxers/README.md)

## Middleware

Through the use of middleware we can wrap requests to applications with commonly run code such as logging, authentication/authorization, and other such tasks.

[Middleware](../../web/middleware/README.md)

## Data Serialization

Before we start building APIs we need to understand how to serialize data. We’ll look at 2 common data formats, as well as ways to customize those formats to match the needs of your API.

[Serializers](../../web/serializers/README.md)

## APIs

By this point we should be able to build fully featured HTML applications in Go, but the fun doesn’t stop there. Let’s turn an eye to building APIs for other applications to consume. We’ll look at two different ways of building APIs (RESTful & HyperMedia), we’ll also look at different ways to handle that age old question of versioning an API.

[Web APIs](../../web/apis/README.md)

## Consuming HTTP APIs

What good is having an API if we can’t consume it? We’ll learn how to use Go to speak with APIs, marshal & unmarshal data, set request headers, and more.

[Consuming APIs](../../web/consuming/README.md)

## Web Sockets

The web is changing and users are expecting fast, dynamic, and interactive web applications. Web Sockets allow for direct two-way communication between the front-end (JavaScript/HTML) and the back-end (Go).

[Web Sockets](../../web/sockets/README.md)

## Authentication

Learn several different techniques and packages for adding authentication to web apps.

[Authentication](../../web/auth/README.md)

## TLS

Sending sensitive data in plain text is a bad idea! Learn about securing your application using TLS.

[TLS](../../web/tls/README.md)

## Shutdown

We can create web servers like professionals now but what about when we need to shut them down? Don't rudely interrupt anyone; shut them down gracefully!

[Shutdown](../../web/shutdown/README.md)
