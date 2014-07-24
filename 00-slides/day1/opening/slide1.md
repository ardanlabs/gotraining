## Purpose Of Go - [FAQ](http://golang.org/doc/faq#What_is_the_purpose_of_the_project)

No major systems language has emerged in over a decade, but over that time the computing landscape has changed tremendously. There are several trends:

* Computers are enormously quicker but software development is not faster.
* Dependency management is a big part of software development today but the “header files” of languages in the C tradition are antithetical to clean dependency analysis—and fast compilation.
* There is a growing rebellion against cumbersome type systems like those of Java and C++, pushing people towards dynamically typed languages such as Python and JavaScript.
* Some fundamental concepts such as garbage collection and parallel computation are not well supported by popular systems languages.
* The emergence of multicore computers has generated worry and confusion.

We believe it's worth trying again with a new language, a concurrent, garbage-collected language with fast compilation. Regarding the points above:

* It is possible to compile a large Go program in a few seconds on a single computer.
* Go provides a model for software construction that makes dependency analysis easy and avoids much of the overhead of C-style include files and libraries.
* Go's type system has no hierarchy, so no time is spent defining the relationships between types. Also, although Go has static types the language attempts to make types feel lighter weight than in typical OO languages.
* Go is fully garbage-collected and provides fundamental support for concurrent execution and communication.
* By its design, Go proposes an approach for the construction of system software on multicore machines.

### [Next](slide2.md)
___
[![GoingGo Training](../../images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../images/ggb_logo.png)](http://www.goinggo.net)