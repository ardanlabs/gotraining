## Text Streams, Redirection, and Pipelining

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

## Example

[Running a subprocess](example1/parent.go)

[Replicating a shell pipeline](example2/pipeline.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
