## Ultimate Web - Introduction to REST
The app is starting to get more complex, at this point we should start talking about design patterns around building web applications, in particular we’ll discuss RESTful design.

### What is REST?

REST stands for Representational State Transfer. It relies on a stateless, client-server, cacheable communications protocol. The HTTP protocol is most commonly used.

REST is an architecture style for designing networked applications. Instead of using complex mechanisms such as CORBA, RPC or SOAP to connect between machines, simple HTTP is used to make calls between machines.

In many ways, the World Wide Web itself, based on HTTP, can be viewed as a REST-based architecture.

RESTful applications use HTTP requests to post data (create and/or update), read data (e.g., make queries), and delete data. Thus, REST uses HTTP for all four CRUD (Create/Read/Update/Delete) operations.

### HTTP Verbs

|HTTP Verb|CRUD           |URL Format     |Common Codes             |
|---      |---            |---            |---                      |
|POST     |Create         |/customers     |200, 201, 404, 409, 422  |
|GET      |Read           |/customers/:id |200, 404                 |
|PUT      |Update/Replace |/customers/:id |200, 204, 404, 422       |
|PATCH    |Update/Modify  |/customers/:id |200, 204, 404, 422       |
|DELETE   |Delete         |/customers/:id |200, 404                 |

[RESTful Server](../../../topics/web/rest/example1/main.go)

#### Exercise

Implement the `PUT`, `PATCH`, and `DELETE` responses in the [RESTful Server](../../../topics/web/rest/example1/main.go).

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
