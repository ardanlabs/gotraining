## Command Invocation

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

## Example

[Invocation](example1/child.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
