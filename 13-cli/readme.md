## Command-Line Interface (CLI) Programming

The command-line is a versatile, efficient, and portable interface that has
seen continuous, ubiquitous use since the mainframe computing era. To this day,
it remains the primary interface for managing server and other production
systems, and is regularly used by computer professionals to complete both
simple and complex tasks with ease.

The capabilities of command-line systems have evolved over time, but the
underlying principles have remained much the same: a command line program
accepts program *arguments*, inherits an *environment*, and operates on input
and output text streams.

## The Command Shell

A program that interprets command descriptions and accordingly arranges for
programs to run is known as a *shell*. It is so named because it encapsulates
the lower level operating system *kernel*; the software that directly manages
system hardware, process invocation and scheduling, and other critical tasks,
but which, unlike a shell, does not any direct user interface.

Shells vary in the syntax they accept and the features they provide, but
typically provide at least: command pipelining, stream redirection, conditional
control flow, pattern matching, job control, and, for interactive shells,
command history.

The [Bourne Shell][sh], released by Bell Labs in 1977, was the first Unix shell
with excellent programming features; it set the standard in shell syntax and
capabilities, and has left a strong legacy of similar shells that are referred
to as "Bourne shell compatible" (or "sh compatible" for short, after the
shell's command name). The following sections are described in terms of Bourne
shell compatible semantics and syntax.

  [sh]: http://en.wikipedia.org/wiki/Bourne_shell

### Command Invocation

A Command consists of a program name optionally followed by space separated
arguments, for example:

    ls
    man ls

Manual pages, accessed via the `man` command, are a core feature of operating
systems following the Unix tradition.

### Text Streams, Redirection, and Pipelining

Every process has its own set of three text streams, collectively known as the
stdio streams: stdin, stdout, and stderr. Programs can be chained together
through a process known as *pipelining*, with the stdout of one process
corresponding to the stdin of the next, and so on. Syntactically, each pair of
commands in a pipeline is separated by a "pipe" character, as shown below:

    # count number of files the current directory
    ls | wc -l

    # list first 10 error lines in a compressed log file
    zcat log.gz | grep -i error | head

Stdio streams can also redirected from or to files. The following two examples
are equivalent: they both lexicographically sort lines from infile, filter out
duplicates, and then write the result to outfile.

    sort < infile | uniq > outfile
    cat infile | sort | uniq | tee outfile

Unless stdin is redirected from a file, the first process in a pipeline will
receive input from the controlling terminal (the keyboard). Likewise, the
stdout of the last process in a pipeline is sent to the controlling terminal,
unless redirected. Stderr is used to communicate information not part of the
normal program output, such as errors; since stderr is not normally redirected,
the error output for all processes in a pipeline are displayed in the terminal.

However, it's often desirable to ignore stderr:

    rm somefile 2>/dev/null

In this case, `2>` means redirect stream #2 (stdin is 0, stdout is 1, and
stderr is 2) to /dev/null, in effect discarding the error output.

### Conditional Control Flow

### Pattern Matching

### Job Control

Pipelines normally run in the *foreground*; there may only be one such pipeline
running in the shell at any given time, and it will receive input and signals
from the controlling terminal.

Pipelines can be alternatively be started in or moved to the *background*.
Background pipelines run normally, except that they do not receive terminal
input, and can only be signalled explicitly by PID or job number. A pipeline
may be suffixed with a single ampersand to run it in the background:

	sleep 5 && echo "wakeup!" &

A foreground pipeline can be suspended by typing `^Z` (Control-Z). Upon doing
so, control will be returned to the shell, and the process can be resumed in
the foreground or background using the `fg` or `bg` commands, respectively.
The `fg` command may also be used to move a running background pipeline to the
foreground.

The `jobs` command lists suspended and background jobs; each of these will have
a *job number*, beginning with 1. The first background job can therefore be
killed with `kill %1`, the second can be moved to the foreground with `fg %2`,
and so on. The percent syntax is known as a *jobspec*.

Finally, the `wait` command can be used wait for one or more background
pipelines, which may be specified using jobspecs. If called without arguments,
wait will block until all active background pipelines have completed.

When a shell (or shell script) exits, all background processes are disowned,
and will continue to run.

Correct use of job control can allow tasks to run concurrently, and make much
more efficient use of computing resources. For example, to build several go
binaries, a good technique would be:

    go build cmd1 &
    go build cmd2 &
    wait

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
