## Go Fuzz

Go-fuzz is a coverage-guided fuzzing solution for testing of Go packages. Fuzzing is mainly applicable to packages that parse complex inputs (both text and binary), and is especially useful for hardening of systems that parse inputs from potentially malicious users (e.g. anything accepted over a network).

## Notes

* Fuzzing allows you to find cases where your code panics.
* Once you identify data inputs that causes panics, code can be corrected and tests created.
* Table tests are an excellent choice for these input data panics.

## Links

https://github.com/dvyukov/go-fuzz  
[go-fuzz github.com/arolek/ase](https://medium.com/@dgryski/go-fuzz-github-com-arolek-ase-3c74d5a3150c#.xvq0ol2zj) - Damian Gryski  
[Go Dynamic Tools](https://www.youtube.com/watch?v=a9xrxRsIbSU) - Dmitry Vyukov  
[DNS parser, meet Go fuzzer](https://blog.cloudflare.com/dns-parser-meet-go-fuzzer) - Filippo Valsorda  
[Fuzzing Beyond Security: Automated Testing with go-fuzz](https://www.youtube.com/watch?v=kOZbFSM7PuI) - Filippo Valsorda  

## Code Review

_When showing a demo, remove the folders under `workdir/corpus` and the `api-fuzz.zip` file._

First thing is to install the Go fuzz tooling:

		go get github.com/dvyukov/go-fuzz/go-fuzz
		go get github.com/dvyukov/go-fuzz/go-fuzz-build

Review the code we want to find problems with and the existing test:

[Code To Fuzz](example1/example1.go) ([Go Playground](http://play.golang.org/p/rDoJiaLOV7))  
[Test For The Code](example1/example1_test.go) ([Go Playground](http://play.golang.org/p/ToGFE_qvJw)) 

Create a corpus file with the initial input data to use and that will be mutated.

[Fuzzing Input Data](workdir/corpus/input.txt) ([Go Playground](http://play.golang.org/p/RkYDgyQsF8))

Create a fuzzing function that takes mutated input and executes the code we care about using it.

[Fuzzing Function](example1/fuzzer.go) ([Go Playground](http://play.golang.org/p/ditM8pfOLm)

Run the `go-fuzz-build` tool against the package to generate the fuzz zip file. The zip file contains all the instrumented binaries go-fuzz is going to use while fuzzing. Any time the source code changes this needs to be re-run.

		go-fuzz-build github.com/ardanlabs/gotraining/topics/fuzzing/example1

Perform the actual fuzzing by running the `go-fuzz` tool and find data inputs that cause panics. Run this until you see an initial crash.

		go-fuzz -bin=./api-fuzz.zip -dup -workdir=workdir/corpus

Review the `crashers` folder under the `workdir/corpus` folders. This contains panic information. You will see an issue when the data passed into the web call is empty. Fix the `Process` function and add the table data to the test.

		{"/process", http.StatusBadRequest, []byte(""), `{"Error":"The Message"}`},

Run the build and start fuzzing.

		rm -rf workdir/crashers and workdir/supressions
		go-fuzz -bin=./api-fuzz.zip -dup -workdir=workdir/corpus

Review the `crashers` folder under the `workdir/corpus` folders. This contains panic information. You will see an issue when value of "0" is used. Fix the `Process` function and add the table data to the test.
		
		{"/process", http.StatusBadRequest, []byte("0"), `{"Error":"The Message"}`},

## Exercises

### Exercise 1

A function named `UnpackUsers` exists and requires fuzzing. A `Fuzz` function and corpus have also been provided. Use the Go Fuzz tooling to find as many problems as you can. In each case fix the bug and add a test table record to validate the test is fixed.

[Exercise](exercises/exercise1)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
