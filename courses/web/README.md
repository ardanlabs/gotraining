## Ultimate Web
Ultimate Web is a 2 day class for any Go developer who wishes to learn how to build robust and well tested HTTP based applications in Go. This class provides an intensive, comprehensive and idiomatic view build Web, SOA, and API applications using Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## HTTP Basics in Go
With a basic understand of how the web and HTTP work, let’s write a simple “Hello World” app in Go. We’ll cover how to start a web server in Go, take in requests, and return responses.

[HTTP Basics in Go](../../topics/web/basics/README.md)

## Testing HTTP in Go
Now that we have some code written, let’s start understanding how to test HTTP Go apps. We’ll look at two different ways of testing HTTP apps.

[HTTP Basics in Go](../../topics/web/testing/README.md)

## HTML Templates
We can now write primitive web apps, as well as test them. Now we need to start adding some meat to it. This section covers generating HTML using Go templates, as well as how to serve up static files, and how to bundle those files in a finished binary.

[HTTP Basics in Go](../../topics/web/templates/README.md)

## POST Requests
HTTP applications don’t just serve content, they also take in content. We’ll branch out of GET requests and start taking in POST requests, processing forms, handling file uploads, and of course, how to test all of this.

[HTTP Basics in Go](../../topics/web/posts/README.md)

## Sessions and Cookies
Managing sessions and cookies is an important part of every web application. Whether it's keeping a user "logged in" or tracking who visits your site, these concepts are essential to learn.

[HTTP Basics in Go](../../topics/web/sessions_cookies/README.md)

## Introduction to REST
The app is starting to get more complex, at this point we should start talking about design patterns around building web applications, in particular we’ll discuss RESTful design.

[HTTP Basics in Go](../../topics/web/rest/README.md)

## Alternative Muxers
The basic muxer in Go has gotten us a long way by this point, but it has it’s limitations. Let’s tour three very different types muxers/routers.

[HTTP Basics in Go](../../topics/web/muxers/README.md)

## Middleware
Through the use of middleware we can wrap requests to applications with commonly run code, such as logging, authentication/authorization, and other such tasks.

[HTTP Basics in Go](../../topics/web/middleware/README.md)

## Data Serialization
Before we start building APIs we need to understand how to serialize data. We’ll look at 3 common data formats, as well as ways to customize those formats to match the needs of your API.

[HTTP Basics in Go](../../topics/web/serializers/README.md)

## APIs
By this point we should be able to build fully featured HTML applications in Go, but the fun doesn’t stop there. Let’s turn an eye to building APIs for other applications to consume. We’ll look at two different ways of building APIs (RESTful & HyperMedia), we’ll also look at different ways to handle that age old question of versioning and API.

[HTTP Basics in Go](../../topics/web/apis/README.md)

## Consuming HTTP APIs
What good is having an API if we can’t consume it? We’ll learn how to use Go to speak with APIs, marshal & unmarshal data, set request headers, and more.

[HTTP Basics in Go](../../topics/web/consuming/README.md)
