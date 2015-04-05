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

The [Bourne Shell][1], released by Bell Labs in 1977, was the first Unix shell
with excellent programming features; it set the standard in shell syntax and
capabilities, and has left a strong legacy of similar shells that are referred
to as "Bourne shell compatible" (or "sh compatible" for short, after the
shell's command name). The following sections are described in terms of Bourne
shell compatible semantics and syntax.

  [1]: http://en.wikipedia.org/wiki/Bourne_shell

### Command Invocation

A Command consists of a program name optionally followed by space separated
arguments, for example:

    ls
    man ls

Manual pages, accessed via the `man` command, are a core feature of operating
systems following the Unix tradition.

Commands inherit an environment, which consists of key-value paired variables.
A parent process may customize the environment each child receives, or may pass
its own environment along without modification. In a shell, environment
variables may be accessed using dollar-sign notation, such as:

    echo "$SHELL"

It's important to quote variables in order to avoid [shell injection][1]
attacks and undefined behavior that can result from special characters and
whitespace in the variable values (shells may re-interpret values after
variable substitution has occurred).

Commands can be invoked with a custom environment by passing key-value pairs
before the command name, such as:

    # add SOMEVAR to the existing environment and display that environment
    SOMEVAR=someval env

    # use the env command as a command runner; in this case it'll just run
    # env again, but with an empty environment
    env -i env

If a variable needs to be in the environment of several commands, it can be
more convenient to "export" it:

    export SOMEVAR=someval
    env

After exporting a variable, all subsequent commands invoked by the shell
process will contain that variable in their environment. Conventionally, all
environment variable names are fully upper-case. As with file and command
names, you should always treat variable names as being case-sensitive.

  [1]: http://en.wikipedia.org/wiki/Code_injection#Shell_injection

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
unless redirected. Even a lone command without explicit redirection is
technically a pipeline, due to its implicit coupling to the terminal.

Stderr is used to communicate information not part of the normal program
output, such as errors; since stderr is not normally redirected, the error
output for all processes in a pipeline are displayed in the terminal.

However, it's often desirable to ignore stderr:

    rm somefile 2>/dev/null

In this case, `2>` means redirect stream #2 (stdin is 0, stdout is 1, and
stderr is 2) to /dev/null, in effect discarding the error output.

### Conditional Control Flow

Commands that complete successfully exit with status zero; all non-zero exit
statuses indicate failure. Conventionally, a status of 1 is used as the generic
program error status, and 2 indicates user error (such as invalid input or
arguments), but any value in the range of 0-127 may be used. On many systems,
statuses 128-255 have special meaning (such as signal reporting) and should not
be set. The exit status of the last process in a pipeline is used as the status
for the entire pipeline.

The success or failure of a pipeline may be used to run additional pipelines:

    rm some-file && echo success || echo failure

Pipelines chained in this way are known as "lists" and have left-to-right
precedence, such that the above will always either echo success or failure, but
never both. Most commonly, AND sequences are chained together so that a series
of tasks can proceed only if every task before it was successful:

    task1 && task2 && task3

If semicolons were used instead (or if each task were on its own line), then
the tasks would run regardless of whether the previous tasks failed.

### Pattern Matching

A fundamental convenience provided by shells is pattern matching of files.
Let's initialize a directory with some files:

    touch a b1 c1 01.txt 02.txt

We can now match against these four files in various ways using patterns.

    # display info for all files ending in '.txt': 01.txt 02.txt
    # the "star" will match any number of characters (including zero)
    ls -l *.txt

    # list all files starting with a lowercase letter: a b1 c1
    # character classes are case-sensitive and match exactly one character
    echo [a-z]*

    # list all files that do _not_ end with a number: a 01.txt 02.txt
    echo *[!0-9]

    # remove all files with names consisting of two characters: b1 c1
    # question marks match exactly one character each
    rm ??

Pattern matching is the default behavior. To instead specify an argument or
filename containing a wildcard character or whitespace, it can be wrapped in
either single or double quotes, or the special character can be escaped with a
backslash, for example:

	# create two new files
	touch filename\ with\ spaces "more spaces"

### Job Control

Command, pipelines, and lists are all jobs, and jobs, by default, run in the
*foreground*; there may only be foreground job running in a shell at any given
time, and it will receive input and signals from the controlling terminal.

Jobs can be alternatively be started in or moved to the *background*.
Background jobs run normally, except that they do not receive terminal input,
and in that state they can only be controlled explicitly by PID or job number.
A job may be suffixed with a single ampersand to run it in the background:

    sleep 5 && echo "wakeup!" &

A foreground job can be suspended by typing `^Z` (Control-Z). Upon doing so,
control will be returned to the shell, and the job can be resumed in the
foreground or background using the `fg` or `bg` commands, respectively.  The
`fg` command may also be used to move a running background job to the
foreground.

The `jobs` command displays suspended and background jobs; each of these will
have a *job number*, beginning with 1. The first background job can therefore
be killed with `kill %1`, the second can be moved to the foreground with `fg
%2`, and so on. The percent syntax is known as a *jobspec*.

Finally, the `wait` command can be used wait for one or more background jobs,
which may be specified using jobspecs. If called without arguments, wait will
block until all active background pipelines have completed.

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
