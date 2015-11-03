## Pattern Matching

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
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
