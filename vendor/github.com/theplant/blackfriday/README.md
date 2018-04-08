Blackfriday
===========

Blackfriday is a [Markdown][1] processor implemented in [Go][2]. It
is paranoid about its input (so you can safely feed it user-supplied
data), it is fast, it supports common extensions (tables, smart
punctuation substitutions, etc.), and it is safe for all utf-8
(unicode) input.

HTML output is currently supported, along with Smartypants
extensions. An experimental LaTeX output engine is also included.

It started as a translation from C of [upskirt][3].


Installation
------------

Assuming you have recent version of Go installed, along with git:

    goinstall github.com/russross/blackfriday

will download, compile, and install the package into
`$GOROOT/src/pkg/github.com/russross/blackfriday`.

For basic usage, it is as simple as getting your input into a byte
slice and calling:

    output := blackfriday.MarkdownBasic(input)

This renders it with no extensions enabled. To get a more useful
feature set, use this instead:

    output := blackfriday.MarkdownCommon(input)

If you want to customize the set of options, first get a renderer
(currently either the HTML or LaTeX output engines), then use it to
call the more general `Markdown` function. For examples, see the
implementations of `MarkdownBasic` and `MarkdownCommon` in
`markdown.go`.

You can also check out `example/main.go` for a more complete example
of how to use it. Run `gomake` in that directory to build a simple
command-line markdown tool:

    cd $GOROOT/src/pkg/github.com/russross/blackfriday/example
    gomake

will build the binary `markdown` in the `example` directory. This is
a statically-linked binary that can be copied to wherever you need
it without worrying about dependencies and library versions.


Features
--------

All features of upskirt are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the `--tidy` option.  Without `--tidy`, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Thread safety**. You can run multiple parsers is different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.


Extensions
----------

In addition to the standard markdown syntax, this package
implements the following extensions:

*   **Intra-word emphasis supression**. The `_` character is
    commonly used inside words when discussing code, so having
    markdown interpret it as an emphasis command is usually the
    wrong thing. Blackfriday lets you treat all emphasis markers as
    normal characters when they occur inside a word.

*   **Tables**. Tables can be created by drawing them in the input
    using a simple syntax:

        Name    | Age
        --------|------
        Bob     | 27
        Alice   | 23

*   **Fenced code blocks**. In addition to the normal 4-space
    indentation to mark code blocks, you can explicitly mark them
    and supply a language (to make syntax highlighting simple). Just
    mark it like this:

        ``` go
        func getTrue() bool {
            return true
        }
        ```

    You can use 3 or more backticks to mark the beginning of the
    block, and the same number to mark the end of the block.

*   **Autolinking**. Blackfriday can find URLs that have not been
    explicitly marked as links and turn them into links.

*   **Strikethrough**. Use two tildes (`~~`) to mark text that
    should be crossed out.

*   **Hard line breaks**. With this extension enabled (it is off by
    default in the `MarkdownBasic` and `MarkdownCommon` convenience
    functions), newlines in the input translate into line breaks in
    the output.

*   **Smart quotes**. Smartypants-style punctuation substitution is
    supported, turning normal double- and single-quote marks into
    curly quotes, etc.

*   **LaTeX-style dash parsing** is an additional option, where `--`
    is translated into `&ndash;`, and `---` is translated into
    `&mdash;`. This differs from most smartypants processors, which
    turn a single hyphen into an ndash and a double hyphen into an
    mdash.

*   **Smart fractions**, where anything that looks like a fraction
    is translated into suitable HTML (instead of just a few special
    cases like most smartypant processors). For example, `4/5`
    becomes `<sup>4</sup>&frasl;<sub>5</sub>`, which renders as
    <sup>4</sup>&frasl;<sub>5</sub>.


LaTeX Output
------------

A rudimentary LaTeX rendering backend is also included. To see an
example of its usage, see `main.go`:

It renders some basic documents, but is only experimental at this
point. In particular, it does not do any inline escaping, so input
that happens to look like LaTeX code will be passed through without
modification.


Todo
----

*   More unit testing
*   Markdown pretty-printer output engine
*   Improve unicode support. It does not understand all unicode
    rules (about what constitutes a letter, a punctuation symbol,
    etc.), so it may fail to detect word boundaries correctly in
    some instances. It is safe on all utf-8 input.


License
-------

Blackfriday is distributed under the Simplified BSD License:

> Copyright © 2011 Russ Ross. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
> 
> The views and conclusions contained in the software and documentation are those of the
> authors and should not be interpreted as representing official policies, either expressed
> or implied, of the copyright holder.


   [1]: http://daringfireball.net/projects/markdown/ "Markdown"
   [2]: http://golang.org/ "Go Language"
   [3]: http://github.com/tanoku/upskirt "Upskirt"
