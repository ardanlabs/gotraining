# Contributing

First of all, if you havn't already, a star would show your support for the project and be very much appreciated!

If you have any suggestions, criticism or bug reports please [create an issue](https://github.com/montanaflynn/stats/issues). 

### Pull Requests

Pull request are always welcome no matter how big or small. Here's an easy way to do it:

1. Fork it and clone your fork
2. Create new branch (`git checkout -b cool-new-thing`)
3. Make the desired changes
4. Ensure tests pass (`go test -cover`)
5. Commit changes (`git commit -am 'Add some cool new thing'`)
6. Push branch (`git push origin cool-new-thing`)
7. Submit Pull Request

#### Advanced 

To make things as seamless as possible please also consider the following steps:

- Write tests to keep 100% code coverage, here's how I check coverage locally without dependencies:

```
go test -coverprofile=coverage.out; go tool cover -html="coverage.out"
```

- Update `README.md` to include new public types or functions in the documentation section.

- Update `examples/main.go` with a simple example of the new feature.

- Run [`gometalinter`](https://github.com/alecthomas/gometalinter) and make your code pass.

- Squash needless commits into single units of work with `git rebase -i new-feature`.

### Makefile

I've included a [Makefile](https://github.com/montanaflynn/stats/blob/master/Makefile) that has a lot of helper targets for common actions such as linting, testing and coverage reporting. 

**Protip: `watch -n 0.5 make check`**
