GitHub Client Exercise
----------------------

Create a project folder in your `$GOPATH` outside of this repository. A good place would be `$GOPATH/src/github.com/yourname/contributors` where `yourname` is your GitHub username.

# Part 1: Just Get it Working

Create a program which can call the GitHub API to get a list of contributors for the `golang/go` repository.

* The API url is https://api.github.com/repos/golang/go/contributors
* Docs for the API in general are https://developer.github.com/v3/
* Docs for the contributors endpoint are https://developer.github.com/v3/repos/#list-contributors
* To get around rate limiting you must generate a personal access token at https://github.com/settings/tokens
* You will be using the `net/http` package to make your request and the `encoding/json` package to decode the response. The docs for each of these packages is online: [net/http](https://golang.org/pkg/net/http), [encoding/json](https://golang.org/pkg/encoding/json).
* To add the authorization header to the request you must first make the request value, add the header, then use a client's `Do` method to execute the request.
* You may have problems calling the GitHub API from within a restricted firewalled network. If you do then open a new terminal in the [githubmock](githubmock) folder and run the program there. Use the url it provides instead of `api.github.com`.

A [template file](template/main.go) is included to get you started.

# Part 2: Refactoring

- Refactor your GitHub API client to a type called `Client` with a method `Contributors`.
- Move your client code from `package main` to another package like `package github`.
- Create a function in package `main` that uses the client. `func
  printContributors(repo string, c *github.Client) int`
- Decouple the `printContributors` function to depend on an interface instead of the concrete type.

# Part 3: Testing

- Add tests for `printContributors`. Create a mock version of the github client
  and pass that in. To capture the ouput you can change the `printContributors`
  function to accept an `io.Writer` where it should print results.
- Add tests for the `github` package using `net/http/httptest.NewServer`.

# Just for Fun
- Add flags to the main package for specifying the repo to pull. Use the [`flag`](https://golang.org/pkg/flag/) package.
- Add a flag to the main package to specify an output file name then encode the results to that file in CSV format. Use the [`encoding/csv`](https://golang.org/pkg/encoding/csv/) package.
- Create a web app that accepts a repo name then shows the contributor list for that repo
