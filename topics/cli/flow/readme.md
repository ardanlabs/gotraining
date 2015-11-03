## Conditional Control Flow

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

## Example

[Setting exit status](example1/parent.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
