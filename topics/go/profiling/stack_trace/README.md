## Stack Traces and Core Dumps

Having some basic skills in debugging Go programs can save any programmer a good amount of time trying to identify problems. I believe in logging as much information as you can, but sometimes a panic occurs and what you logged is not enough. Understanding the information in a stack trace can sometimes mean the difference between finding the bug now or needing to add more logging and waiting for it to happen again. We can also stop a running program and get Core Dump which also generates a stack trace.

## Notes

* Stack traces are an important tool in debugging an application.
* The runtime should never panic so the trace is everything.
* You can see every goroutine and the call stack for each routine.
* You can see every value passed into each function on the stack.
* You can generate core dumps and use these same techniques.

## Stack Traces

These two programs called the built-in function `panic` to produce a stack trace. Stack traces show not only the call stack from the line of code that caused the panic. They also show the values that were passed into each function.

### Example 1

Build and run the program.

    $ go build
    $ ./example1

Review the stack trace.

    // Stack Trace
    goroutine 1 [running]:
    main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa, 0x0, 0xc000054778)
        stack_trace/example1/example1.go:13 +0x39
    main.main()
        stack_trace/example1/example1.go:8 +0x85

    // Declaration
    main.example(slice []string, str string, i int) error

    // Call
    main.example(make([]string, 2, 4), "hello", 10)

    // Values (0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa, 0x0, 0xc000054778)
    Slice Value:      0xc000042748, 0x2, 0x4
    String Value:     0x106abae, 0x5
    Integer Value:    0xa
    Return Arguments: 0x0, 0xc000054778

Use `go build -gcflags -S` to map the PC offset values, +0x39 and +0x72 for
each function call.

### Example 2

Build and run the program.

    $ go build
    $ ./example2

Review the stack trace.

    // Stack Trace
    goroutine 1 [running]:
    main.example(0xc019010001, 0x1064ee0, 0xc000076058)
        stack_trace/example2/example2.go:13 +0x39
    main.main()
        stack_trace/example2/example2.go:8 +0x29

    // Declaration
    main.example(b1, b2, b3 bool, i uint8)

    // Call
    main.example(true, false, true, 25)

    // Word value (0xc019010001)
    Bits    Binary      Hex   Value
    00-07   0000 0001   01    true
    08-15   0000 0000   00    false
    16-23   0000 0001   01    true
    24-31   0001 1001   19    25

    Return Arguments: 0x1064ee0, 0xc00007605

Use `go build -gcflags -S` to map the PC offset values, +0x39 and +0x29 for
each function call.

### Code Review

[Review Stack Trace](example1/example1.go) ([Go Playground](https://play.golang.org/p/k18FqfsuHdU))  
[Packing](example2/example2.go) ([Go Playground](https://play.golang.org/p/WhGxuICFhLu))  

### Links

[Stack traces in Go](https://www.ardanlabs.com/blog/2015/01/stack-traces-in-go.html)  

## Core Dumps

You can generate a core dump of any running Go program by issuing a SIGQUIT to the program. You can do this by pressing (Ctrl+\\) on your keyboard.

### Generating a Core Dump

Build and run the example program.

    $ go build
    $ ./godebug

Put some load of the web application.

    $ hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"

Issue a signal quit.

    Ctrl+\

Review the dump.

    2017/05/31 06:24:24 listener : Started : Listening on: http://localhost:4000 : Leak[false]
    ^\SIGQUIT: quit
    PC=0x10564cb m=0 sigcode=0

    goroutine 0 [idle]:
    runtime.mach_semaphore_wait(0xf03, 0x13e2040, 0x7fff5fbff860, 0x1030cda, 0x13e1c30, 0x13e2040, 0x7fff5fbff810, 0x1050ba3, 0xffffffffffffffff, 0x1, ...)
        /usr/local/go/src/runtime/sys_darwin_amd64.s:415 +0xb
    runtime.semasleep1(0xffffffffffffffff, 0x1)
        /usr/local/go/src/runtime/os_darwin.go:413 +0x4b

    ...

    rax    0xe
    rbx    0x13e23a0
    rcx    0x7fff5fbff7b0

Get a larger crash dump by running the program using the GOTRACEBACK env variable.

    $ GOTRACEBACK=crash ./crash

Use `gcore` to write a dump of the running program.

    $ sudo gcore -o core.txt PID

Use Delve to review the dump.

    $ dlv core ./godebug core.txt

    (dlv) bt
    0  0x0000000000457774 in runtime.raise
        at /usr/lib/go/src/runtime/sys_linux_amd64.s:110
    1  0x000000000043f7fb in runtime.dieFromSignal

    (dlv) ls
    > runtime.raise() /usr/lib/go/src/runtime/sys_linux_amd64.s:110 (PC: 0x457774)
    105:		SYSCALL
    106:		MOVL	AX, DI	// arg 1 tid
    107:		MOVL	sig+0(FP), SI	// arg 2
    108:		MOVL	$200, AX	// syscall - tkill
    109:		SYSCALL
    => 110:		RET

### Links

[Debugging Go core dumps](https://rakyll.org/coredumps/) - JBD    

### Code Review

[Core Dumps](example3/example3.go) ([Go Playground](https://play.golang.org/p/rPVBbcQhFeX))  
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
