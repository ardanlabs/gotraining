//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
// Unit tests for block parsing
//

package blackfriday

import (
	"testing"
)

func runMarkdownBlock(input string, extensions int) string {
	htmlFlags := 0
	htmlFlags |= HTML_USE_XHTML

	renderer := HtmlRenderer(htmlFlags, "", "")

	return string(Markdown([]byte(input), renderer, extensions))
}

func doTestsBlock(t *testing.T, tests []string, extensions int) {
	// catch and report panics
	var candidate string
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("\npanic while processing [%#v]\n", candidate)
		}
	}()

	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		candidate = input
		expected := tests[i+1]
		actual := runMarkdownBlock(candidate, extensions)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}

		// now test every substring to stress test bounds checking
		if !testing.Short() {
			for start := 0; start < len(input); start++ {
				for end := start + 1; end <= len(input); end++ {
					candidate = input[start:end]
					_ = runMarkdownBlock(candidate, extensions)
				}
			}
		}
	}
}

func TestPrefixHeaderNoExtensions(t *testing.T) {
	var tests = []string{
		"#\n",
		"<h1></h1>\n",

		"# \n",
		"<h1></h1>\n",

		`#

![](/5018d345558fbe46c4000001/537f27fc18c8bca41a000010/file/5422741b6c08982b400000be/m/undefined "undefined")

He Orders a beer.

![](/5018d345558fbe46c4000001/537f27fc18c8bca41a000010/file/542274696c08982b400000da/m/undefined "undefined")

# He Orders 0 beers.`,

		`<h1></h1>

<p><img src="/5018d345558fbe46c4000001/537f27fc18c8bca41a000010/file/5422741b6c08982b400000be/m/undefined" alt="" title="undefined" />
</p>

<p>He Orders a beer.</p>

<p><img src="/5018d345558fbe46c4000001/537f27fc18c8bca41a000010/file/542274696c08982b400000da/m/undefined" alt="" title="undefined" />
</p>

<h1>He Orders 0 beers.</h1>
`,

		"# Header 1\n",
		"<h1>Header 1</h1>\n",

		"## Header 2\n",
		"<h2>Header 2</h2>\n",

		"### Header 3\n",
		"<h3>Header 3</h3>\n",

		"#### Header 4\n",
		"<h4>Header 4</h4>\n",

		"##### Header 5\n",
		"<h5>Header 5</h5>\n",

		"###### Header 6\n",
		"<h6>Header 6</h6>\n",

		"####### Header 7\n",
		"<h6># Header 7</h6>\n",

		"#Header 1\n",
		"<h1>Header 1</h1>\n",

		"##Header 2\n",
		"<h2>Header 2</h2>\n",

		"###Header 3\n",
		"<h3>Header 3</h3>\n",

		"####Header 4\n",
		"<h4>Header 4</h4>\n",

		"#####Header 5\n",
		"<h5>Header 5</h5>\n",

		"######Header 6\n",
		"<h6>Header 6</h6>\n",

		"#######Header 7\n",
		"<h6>#Header 7</h6>\n",

		"Hello\n# Header 1\nGoodbye\n",
		"<p>Hello</p>\n\n<h1>Header 1</h1>\n\n<p>Goodbye</p>\n",

		"* List\n# Header\n* List\n",
		"<ul>\n<li><p>List</p>\n\n<h1>Header</h1></li>\n\n<li><p>List</p></li>\n</ul>\n",

		"* List\n#Header\n* List\n",
		"<ul>\n<li><p>List</p>\n\n<h1>Header</h1></li>\n\n<li><p>List</p></li>\n</ul>\n",

		"*   List\n    * Nested list\n    # Nested header\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li><p>Nested list</p>\n\n" +
			"<h1>Nested header</h1></li>\n</ul></li>\n</ul>\n",
	}
	doTestsBlock(t, tests, 0)
}

func TestNoListItemBlockExtension(t *testing.T) {
	var tests = []string{

		`# Title
* a point
  it worked?

* second point

* third point`,

		`<h1>Title</h1>

<ul>
<li>a point
it worked?</li>

<li>second point</li>

<li>third point</li>
</ul>
`,
	}
	doTestsBlock(t, tests, 0|EXTENSION_NO_LIST_ITEM_BLOCK)
}

func TestPrefixHeaderSpaceExtension(t *testing.T) {
	var tests = []string{
		"# Header 1\n",
		"<h1>Header 1</h1>\n",

		"## Header 2\n",
		"<h2>Header 2</h2>\n",

		"### Header 3\n",
		"<h3>Header 3</h3>\n",

		"#### Header 4\n",
		"<h4>Header 4</h4>\n",

		"##### Header 5\n",
		"<h5>Header 5</h5>\n",

		"###### Header 6\n",
		"<h6>Header 6</h6>\n",

		"####### Header 7\n",
		"<p>####### Header 7</p>\n",

		"#Header 1\n",
		"<p>#Header 1</p>\n",

		"##Header 2\n",
		"<p>##Header 2</p>\n",

		"###Header 3\n",
		"<p>###Header 3</p>\n",

		"####Header 4\n",
		"<p>####Header 4</p>\n",

		"#####Header 5\n",
		"<p>#####Header 5</p>\n",

		"######Header 6\n",
		"<p>######Header 6</p>\n",

		"#######Header 7\n",
		"<p>#######Header 7</p>\n",

		"Hello\n# Header 1\nGoodbye\n",
		"<p>Hello</p>\n\n<h1>Header 1</h1>\n\n<p>Goodbye</p>\n",

		"* List\n# Header\n* List\n",
		"<ul>\n<li><p>List</p>\n\n<h1>Header</h1></li>\n\n<li><p>List</p></li>\n</ul>\n",

		"* List\n#Header\n* List\n",
		"<ul>\n<li>List\n#Header</li>\n<li>List</li>\n</ul>\n",

		"*   List\n    * Nested list\n    # Nested header\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li><p>Nested list</p>\n\n" +
			"<h1>Nested header</h1></li>\n</ul></li>\n</ul>\n",
	}
	doTestsBlock(t, tests, EXTENSION_SPACE_HEADERS)
}

func TestUnderlineHeaders(t *testing.T) {
	var tests = []string{
		"Header 1\n========\n",
		"<h1>Header 1</h1>\n",

		"Header 2\n--------\n",
		"<h2>Header 2</h2>\n",

		"A\n=\n",
		"<h1>A</h1>\n",

		"B\n-\n",
		"<h2>B</h2>\n",

		"Paragraph\nHeader\n=\n",
		"<p>Paragraph</p>\n\n<h1>Header</h1>\n",

		"Header\n===\nParagraph\n",
		"<h1>Header</h1>\n\n<p>Paragraph</p>\n",

		"Header\n===\nAnother header\n---\n",
		"<h1>Header</h1>\n\n<h2>Another header</h2>\n",

		"   Header\n======\n",
		"<h1>Header</h1>\n",

		"    Code\n========\n",
		"<pre><code>Code\n</code></pre>\n\n<p>========</p>\n",

		"Header with *inline*\n=====\n",
		"<h1>Header with <em>inline</em></h1>\n",

		"*   List\n    * Sublist\n    Not a header\n    ------\n",
		"<ul>\n<li>List\n\n<ul>\n<li>Sublist\nNot a header\n------</li>\n</ul></li>\n</ul>\n",

		"Paragraph\n\n\n\n\nHeader\n===\n",
		"<p>Paragraph</p>\n\n<h1>Header</h1>\n",

		"Trailing space \n====        \n\n",
		"<h1>Trailing space</h1>\n",

		"Trailing spaces\n====        \n\n",
		"<h1>Trailing spaces</h1>\n",

		"Double underline\n=====\n=====\n",
		"<h1>Double underline</h1>\n\n<p>=====</p>\n",
	}
	doTestsBlock(t, tests, 0)
}

func TestHorizontalRule(t *testing.T) {
	var tests = []string{
		"-\n",
		"<p>-</p>\n",

		"--\n",
		"<p>--</p>\n",

		"---\n",
		"<hr />\n",

		"----\n",
		"<hr />\n",

		"*\n",
		"<p>*</p>\n",

		"**\n",
		"<p>**</p>\n",

		"***\n",
		"<hr />\n",

		"****\n",
		"<hr />\n",

		"_\n",
		"<p>_</p>\n",

		"__\n",
		"<p>__</p>\n",

		"___\n",
		"<hr />\n",

		"____\n",
		"<hr />\n",

		"-*-\n",
		"<p>-*-</p>\n",

		"- - -\n",
		"<hr />\n",

		"* * *\n",
		"<hr />\n",

		"_ _ _\n",
		"<hr />\n",

		"-----*\n",
		"<p>-----*</p>\n",

		"   ------   \n",
		"<hr />\n",

		"Hello\n***\n",
		"<p>Hello</p>\n\n<hr />\n",

		"---\n***\n___\n",
		"<hr />\n\n<hr />\n\n<hr />\n",
	}
	doTestsBlock(t, tests, 0)
}

func TestUnorderedList(t *testing.T) {
	var tests = []string{
		"* Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"* Yin\n* Yang\n",
		"<ul>\n<li>Yin</li>\n<li>Yang</li>\n</ul>\n",

		"* Ting\n* Bong\n* Goo\n",
		"<ul>\n<li>Ting</li>\n<li>Bong</li>\n<li>Goo</li>\n</ul>\n",

		"* Yin\n\n* Yang\n",
		"<ul>\n<li><p>Yin</p></li>\n\n<li><p>Yang</p></li>\n</ul>\n",

		"* Ting\n\n* Bong\n* Goo\n",
		"<ul>\n<li><p>Ting</p></li>\n\n<li><p>Bong</p></li>\n\n<li><p>Goo</p></li>\n</ul>\n",

		"+ Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"+ Yin\n+ Yang\n",
		"<ul>\n<li>Yin</li>\n<li>Yang</li>\n</ul>\n",

		"+ Ting\n+ Bong\n+ Goo\n",
		"<ul>\n<li>Ting</li>\n<li>Bong</li>\n<li>Goo</li>\n</ul>\n",

		"+ Yin\n\n+ Yang\n",
		"<ul>\n<li><p>Yin</p></li>\n\n<li><p>Yang</p></li>\n</ul>\n",

		"+ Ting\n\n+ Bong\n+ Goo\n",
		"<ul>\n<li><p>Ting</p></li>\n\n<li><p>Bong</p></li>\n\n<li><p>Goo</p></li>\n</ul>\n",

		"- Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"- Yin\n- Yang\n",
		"<ul>\n<li>Yin</li>\n<li>Yang</li>\n</ul>\n",

		"- Ting\n- Bong\n- Goo\n",
		"<ul>\n<li>Ting</li>\n<li>Bong</li>\n<li>Goo</li>\n</ul>\n",

		"- Yin\n\n- Yang\n",
		"<ul>\n<li><p>Yin</p></li>\n\n<li><p>Yang</p></li>\n</ul>\n",

		"- Ting\n\n- Bong\n- Goo\n",
		"<ul>\n<li><p>Ting</p></li>\n\n<li><p>Bong</p></li>\n\n<li><p>Goo</p></li>\n</ul>\n",

		"*   Hello \n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"*   Hello \n    Next line \n",
		"<ul>\n<li>Hello\nNext line</li>\n</ul>\n",

		"Paragraph\n\n* Linebreak\n",
		"<p>Paragraph</p>\n\n<ul>\n<li>Linebreak</li>\n</ul>\n",

		"*   List\n    * Nested list\n",
		"<ul>\n<li>List\n\n<ul>\n<li>Nested list</li>\n</ul></li>\n</ul>\n",

		"*   List\n\n    * Nested list\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li>Nested list</li>\n</ul></li>\n</ul>\n",

		"*   List\n    Second line\n\n    + Nested\n",
		"<ul>\n<li><p>List\nSecond line</p>\n\n<ul>\n<li>Nested</li>\n</ul></li>\n</ul>\n",

		"*   List\n    + Nested\n\n    Continued\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li>Nested</li>\n</ul>\n\n<p>Continued</p></li>\n</ul>\n",

		"*   List\n   * shallow indent\n",
		"<ul>\n<li>List\n\n<ul>\n<li>shallow indent</li>\n</ul></li>\n</ul>\n",

		"* List\n" +
			" * shallow indent\n" +
			"  * part of second list\n" +
			"   * still second\n" +
			"    * almost there\n" +
			"     * third level\n",
		"<ul>\n" +
			"<li>List\n\n" +
			"<ul>\n" +
			"<li>shallow indent</li>\n" +
			"<li>part of second list</li>\n" +
			"<li>still second</li>\n" +
			"<li>almost there\n\n" +
			"<ul>\n" +
			"<li>third level</li>\n" +
			"</ul></li>\n" +
			"</ul></li>\n" +
			"</ul>\n",

		"* List\n        extra indent, same paragraph\n",
		"<ul>\n<li>List\n    extra indent, same paragraph</li>\n</ul>\n",

		"* List\n\n    >indent quote\n",
		"<ul>\n<li><p>List</p>\n\n<blockquote>\n<p>indent quote</p>\n</blockquote></li>\n</ul>\n",

		"* List\n\n        code block\n",
		"<ul>\n<li><p>List</p>\n\n<pre><code>code block\n</code></pre></li>\n</ul>\n",

		"* List\n\n          code block with spaces\n",
		"<ul>\n<li><p>List</p>\n\n<pre><code>  code block with spaces\n</code></pre></li>\n</ul>\n",

		"* List\n\n    * sublist\n\n    normal text\n\n    * another sublist\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li>sublist</li>\n</ul>\n\n<p>normal text</p>\n\n<ul>\n<li>another sublist</li>\n</ul></li>\n</ul>\n",
	}

	var tests1 = append(tests, []string{
		"*Hello\n",
		"<p>*Hello</p>\n",

		"Paragraph\n* No linebreak\n",
		"<p>Paragraph\n* No linebreak</p>\n",
	}...)
	doTestsBlock(t, tests1, 0)

	var tests2 = append(tests, []string{
		"*Hello\n",
		"<p>*Hello</p>\n",

		"Paragraph\n* No linebreak\n",
		"<p>Paragraph</p>\n\n<ul>\n<li>No linebreak</li>\n</ul>\n",
	}...)
	doTestsBlock(t, tests2, EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK)

	var tests3 = append(tests, []string{
		"*Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"-Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"+Hello\n",
		"<ul>\n<li>Hello</li>\n</ul>\n",

		"*Hello*\n",
		"<p><em>Hello</em></p>\n",

		"-Hello-\n",
		"<ul>\n<li>Hello-</li>\n</ul>\n",

		"+Hello+\n",
		"<ul>\n<li>Hello+</li>\n</ul>\n",

		"Paragraph\n* No linebreak\n",
		"<p>Paragraph\n* No linebreak</p>\n",

		"Paragraph\n\n*Linebreak\n",
		"<p>Paragraph</p>\n\n<ul>\n<li>Linebreak</li>\n</ul>\n",

		"-*-\n",
		"<ul>\n<li>*-</li>\n</ul>\n",

		"*-*\n",
		"<p><em>-</em></p>\n",

		"*at the* beginning\n",
		"<p><em>at the</em> beginning</p>\n",

		"*try two* in *one line*\n",
		"<p><em>try two</em> in <em>one line</em></p>\n",

		"*__triple emphasis__*\n",
		"<p><em><strong>triple emphasis</strong></em></p>\n",

		"*improper **nesting* is** bad\n",
		"<p><em>improper </em><em>nesting</em> is** bad</p>\n",

		"-http://foo.com/\n",
		"<ul>\n<li><a href=\"http://foo.com/\">http://foo.com/</a></li>\n</ul>\n",

		"*http://foo.com/\n",
		"<ul>\n<li><a href=\"http://foo.com/\">http://foo.com/</a></li>\n</ul>\n",

		"*List\n" +
			" *shallow indent\n" +
			"  *part of second list\n" +
			"   *still second\n" +
			"    *almost there\n" +
			"     *third level\n",
		"<ul>\n" +
			"<li>List\n\n" +
			"<ul>\n" +
			"<li>shallow indent</li>\n" +
			"<li>part of second list</li>\n" +
			"<li>still second</li>\n" +
			"<li>almost there\n\n" +
			"<ul>\n" +
			"<li>third level</li>\n" +
			"</ul></li>\n" +
			"</ul></li>\n" +
			"</ul>\n",

		"*List\n        extra indent, same paragraph\n",
		"<ul>\n<li>List\n    extra indent, same paragraph</li>\n</ul>\n",

		"*List\n\n        code block\n",
		"<ul>\n<li><p>List</p>\n\n<pre><code>code block\n</code></pre></li>\n</ul>\n",

		"*List\n\n          code block with spaces\n",
		"<ul>\n<li><p>List</p>\n\n<pre><code>  code block with spaces\n</code></pre></li>\n</ul>\n",

		"*List\n\n    *sublist\n\n    normal text\n\n    *another sublist\n",
		"<ul>\n<li><p>List</p>\n\n<ul>\n<li>sublist</li>\n</ul>\n\n<p>normal text</p>\n\n<ul>\n<li>another sublist</li>\n</ul></li>\n</ul>\n",
	}...)
	doTestsBlock(t, tests3, EXTENSION_NO_SPACE_LISTS|EXTENSION_AUTOLINK)

	var tests4 = append(tests, []string{
		"* List\n\n 1 space indent paragraph\n",
		"<ul>\n<li><p>List</p>\n\n<p>1 space indent paragraph</p></li>\n</ul>\n",

		"* List\n\n  2 spaces indent paragraph\n",
		"<ul>\n<li><p>List</p>\n\n<p>2 spaces indent paragraph</p></li>\n</ul>\n",

		"* List\n\n   3 spaces indent paragraph\n",
		"<ul>\n<li><p>List</p>\n\n<p>3 spaces indent paragraph</p></li>\n</ul>\n",

		"* List\n\n >1 space indent quote\n",
		"<ul>\n<li><p>List</p>\n\n<blockquote>\n<p>1 space indent quote</p>\n</blockquote></li>\n</ul>\n",

		"* List\n\n  >2 spaces indent quote\n",
		"<ul>\n<li><p>List</p>\n\n<blockquote>\n<p>2 spaces indent quote</p>\n</blockquote></li>\n</ul>\n",

		"* List\n\n   >3 spaces indent quote\n",
		"<ul>\n<li><p>List</p>\n\n<blockquote>\n<p>3 spaces indent quote</p>\n</blockquote></li>\n</ul>\n",
	}...)
	doTestsBlock(t, tests4, EXTENSION_ONE_SPACE_INDENT)

	var tests5 = append(tests, []string{
		"– Item1\n\n – Item2\n\n",
		"<ul>\n<li><p>Item1</p>\n\n<ul>\n<li>Item2</li>\n</ul></li>\n</ul>\n",
	}...)
	doTestsBlock(t, tests5, EXTENSION_UNICODE_LIST_ITEM)
}

func TestOrderedList(t *testing.T) {
	var tests = []string{
		"1. Hello\n",
		"<ol>\n<li>Hello</li>\n</ol>\n",

		"1. Yin\n2. Yang\n",
		"<ol>\n<li>Yin</li>\n<li>Yang</li>\n</ol>\n",

		"1. Ting\n2. Bong\n3. Goo\n",
		"<ol>\n<li>Ting</li>\n<li>Bong</li>\n<li>Goo</li>\n</ol>\n",

		"1. Yin\n\n2. Yang\n",
		"<ol>\n<li><p>Yin</p></li>\n\n<li><p>Yang</p></li>\n</ol>\n",

		"1. Ting\n\n2. Bong\n3. Goo\n",
		"<ol>\n<li><p>Ting</p></li>\n\n<li><p>Bong</p></li>\n\n<li><p>Goo</p></li>\n</ol>\n",

		"1 Hello\n",
		"<p>1 Hello</p>\n",

		"1.  Hello \n",
		"<ol>\n<li>Hello</li>\n</ol>\n",

		"1.  Hello \n    Next line \n",
		"<ol>\n<li>Hello\nNext line</li>\n</ol>\n",

		"Paragraph\n\n1. Linebreak\n",
		"<p>Paragraph</p>\n\n<ol>\n<li>Linebreak</li>\n</ol>\n",

		"1.  List\n    1. Nested list\n",
		"<ol>\n<li>List\n\n<ol>\n<li>Nested list</li>\n</ol></li>\n</ol>\n",

		"1.  List\n\n    1. Nested list\n",
		"<ol>\n<li><p>List</p>\n\n<ol>\n<li>Nested list</li>\n</ol></li>\n</ol>\n",

		"1.  List\n    Second line\n\n    1. Nested\n",
		"<ol>\n<li><p>List\nSecond line</p>\n\n<ol>\n<li>Nested</li>\n</ol></li>\n</ol>\n",

		"1.  List\n    1. Nested\n\n    Continued\n",
		"<ol>\n<li><p>List</p>\n\n<ol>\n<li>Nested</li>\n</ol>\n\n<p>Continued</p></li>\n</ol>\n",

		"1.  List\n   1. shallow indent\n",
		"<ol>\n<li>List\n\n<ol>\n<li>shallow indent</li>\n</ol></li>\n</ol>\n",

		"1. List\n" +
			" 1. shallow indent\n" +
			"  2. part of second list\n" +
			"   3. still second\n" +
			"    4. almost there\n" +
			"     1. third level\n",
		"<ol>\n" +
			"<li>List\n\n" +
			"<ol>\n" +
			"<li>shallow indent</li>\n" +
			"<li>part of second list</li>\n" +
			"<li>still second</li>\n" +
			"<li>almost there\n\n" +
			"<ol>\n" +
			"<li>third level</li>\n" +
			"</ol></li>\n" +
			"</ol></li>\n" +
			"</ol>\n",

		"1. List\n        extra indent, same paragraph\n",
		"<ol>\n<li>List\n    extra indent, same paragraph</li>\n</ol>\n",

		"1. List\n\n        code block\n",
		"<ol>\n<li><p>List</p>\n\n<pre><code>code block\n</code></pre></li>\n</ol>\n",

		"1. List\n\n          code block with spaces\n",
		"<ol>\n<li><p>List</p>\n\n<pre><code>  code block with spaces\n</code></pre></li>\n</ol>\n",

		"1. List\n    * Mixted list\n",
		"<ol>\n<li>List\n\n<ul>\n<li>Mixted list</li>\n</ul></li>\n</ol>\n",

		"1. List\n * Mixed list\n",
		"<ol>\n<li>List\n\n<ul>\n<li>Mixed list</li>\n</ul></li>\n</ol>\n",

		"* Start with unordered\n 1. Ordered\n",
		"<ul>\n<li>Start with unordered\n\n<ol>\n<li>Ordered</li>\n</ol></li>\n</ul>\n",

		"* Start with unordered\n    1. Ordered\n",
		"<ul>\n<li>Start with unordered\n\n<ol>\n<li>Ordered</li>\n</ol></li>\n</ul>\n",

		"1. numbers\n1. are ignored\n",
		"<ol>\n<li>numbers</li>\n<li>are ignored</li>\n</ol>\n",
	}
	var tests1 = append(tests, []string{
		"1.Hello\n",
		"<p>1.Hello</p>\n",

		"Paragraph\n1. No linebreak\n",
		"<p>Paragraph\n1. No linebreak</p>\n",
	}...)
	doTestsBlock(t, tests1, 0)

	var tests2 = append(tests, []string{
		"1.Hello\n",
		"<p>1.Hello</p>\n",

		"Paragraph\n1. No linebreak\n",
		"<p>Paragraph</p>\n\n<ol>\n<li>No linebreak</li>\n</ol>\n",
	}...)
	doTestsBlock(t, tests2, EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK)

	var tests3 = append(tests, []string{
		"1.Hello\n",
		"<ol>\n<li>Hello</li>\n</ol>\n",

		"1.Yin\n2.Yang\n",
		"<ol>\n<li>Yin</li>\n<li>Yang</li>\n</ol>\n",

		"1.Ting\n2.Bong\n3.Goo\n",
		"<ol>\n<li>Ting</li>\n<li>Bong</li>\n<li>Goo</li>\n</ol>\n",

		"1.Yin\n\n2.Yang\n",
		"<ol>\n<li><p>Yin</p></li>\n\n<li><p>Yang</p></li>\n</ol>\n",

		"1.Ting\n\n2.Bong\n3. Goo\n",
		"<ol>\n<li><p>Ting</p></li>\n\n<li><p>Bong</p></li>\n\n<li><p>Goo</p></li>\n</ol>\n",

		"1.Hello \n    Next line \n",
		"<ol>\n<li>Hello\nNext line</li>\n</ol>\n",

		"Paragraph\n\n1.Linebreak\n",
		"<p>Paragraph</p>\n\n<ol>\n<li>Linebreak</li>\n</ol>\n",

		"1.List\n    1.Nested list\n",
		"<ol>\n<li>List\n\n<ol>\n<li>Nested list</li>\n</ol></li>\n</ol>\n",

		"1.List\n\n    1.Nested list\n",
		"<ol>\n<li><p>List</p>\n\n<ol>\n<li>Nested list</li>\n</ol></li>\n</ol>\n",

		"1.List\n    Second line\n\n    1.Nested\n",
		"<ol>\n<li><p>List\nSecond line</p>\n\n<ol>\n<li>Nested</li>\n</ol></li>\n</ol>\n",

		"1.List\n    1.Nested\n\n    Continued\n",
		"<ol>\n<li><p>List</p>\n\n<ol>\n<li>Nested</li>\n</ol>\n\n<p>Continued</p></li>\n</ol>\n",

		"1.List\n   1.shallow indent\n",
		"<ol>\n<li>List\n\n<ol>\n<li>shallow indent</li>\n</ol></li>\n</ol>\n",

		"1.List\n" +
			" 1.shallow indent\n" +
			"  2.part of second list\n" +
			"   3.still second\n" +
			"    4.almost there\n" +
			"     1.third level\n",
		"<ol>\n" +
			"<li>List\n\n" +
			"<ol>\n" +
			"<li>shallow indent</li>\n" +
			"<li>part of second list</li>\n" +
			"<li>still second</li>\n" +
			"<li>almost there\n\n" +
			"<ol>\n" +
			"<li>third level</li>\n" +
			"</ol></li>\n" +
			"</ol></li>\n" +
			"</ol>\n",

		"1.List\n        extra indent, same paragraph\n",
		"<ol>\n<li>List\n    extra indent, same paragraph</li>\n</ol>\n",

		"1.List\n\n        code block\n",
		"<ol>\n<li><p>List</p>\n\n<pre><code>code block\n</code></pre></li>\n</ol>\n",

		"1.List\n\n          code block with spaces\n",
		"<ol>\n<li><p>List</p>\n\n<pre><code>  code block with spaces\n</code></pre></li>\n</ol>\n",

		"1.List\n    *Mixted list\n",
		"<ol>\n<li>List\n\n<ul>\n<li>Mixted list</li>\n</ul></li>\n</ol>\n",

		"1.List\n *Mixed list\n",
		"<ol>\n<li>List\n\n<ul>\n<li>Mixed list</li>\n</ul></li>\n</ol>\n",

		"*Start with unordered\n 1.Ordered\n",
		"<ul>\n<li>Start with unordered\n\n<ol>\n<li>Ordered</li>\n</ol></li>\n</ul>\n",

		"*Start with unordered\n    1.Ordered\n",
		"<ul>\n<li>Start with unordered\n\n<ol>\n<li>Ordered</li>\n</ol></li>\n</ul>\n",

		"1.numbers\n1.are ignored\n",
		"<ol>\n<li>numbers</li>\n<li>are ignored</li>\n</ol>\n",

		"Qortex costs\n\n1.9 billion!",
		"<p>Qortex costs</p>\n\n<p>1.9 billion!</p>\n",

		"1.",
		"<p>1.</p>\n",
	}...)
	doTestsBlock(t, tests3, EXTENSION_NO_SPACE_LISTS)

	var tests4 = append(tests, []string{
		"1. List\n\n 1 space indent paragraph\n",
		"<ol>\n<li><p>List</p>\n\n<p>1 space indent paragraph</p></li>\n</ol>\n",

		"1. List\n\n  2 spaces indent paragraph\n",
		"<ol>\n<li><p>List</p>\n\n<p>2 spaces indent paragraph</p></li>\n</ol>\n",

		"1. List\n\n   3 spaces indent paragraph\n",
		"<ol>\n<li><p>List</p>\n\n<p>3 spaces indent paragraph</p></li>\n</ol>\n",

		"1. List\n\n >1 space indent quote\n",
		"<ol>\n<li><p>List</p>\n\n<blockquote>\n<p>1 space indent quote</p>\n</blockquote></li>\n</ol>\n",

		"1. List\n\n  >2 spaces indent quote\n",
		"<ol>\n<li><p>List</p>\n\n<blockquote>\n<p>2 spaces indent quote</p>\n</blockquote></li>\n</ol>\n",

		"1. List\n\n   >3 spaces indent quote\n",
		"<ol>\n<li><p>List</p>\n\n<blockquote>\n<p>3 spaces indent quote</p>\n</blockquote></li>\n</ol>\n",
	}...)
	doTestsBlock(t, tests4, EXTENSION_ONE_SPACE_INDENT)
}

func TestPreformattedHtml(t *testing.T) {
	var tests = []string{
		"<div></div>\n",
		"<div></div>\n",

		"<div>\n</div>\n",
		"<div>\n</div>\n",

		"<div>\n</div>\nParagraph\n",
		"<p><div>\n</div>\nParagraph</p>\n",

		"<div class=\"foo\">\n</div>\n",
		"<div class=\"foo\">\n</div>\n",

		"<div>\nAnything here\n</div>\n",
		"<div>\nAnything here\n</div>\n",

		"<div>\n  Anything here\n</div>\n",
		"<div>\n  Anything here\n</div>\n",

		"<div>\nAnything here\n  </div>\n",
		"<div>\nAnything here\n  </div>\n",

		"<div>\nThis is *not* &proceessed\n</div>\n",
		"<div>\nThis is *not* &proceessed\n</div>\n",

		"<faketag>\n  Something\n</faketag>\n",
		"<p><faketag>\n  Something\n</faketag></p>\n",

		"<div>\n  Something here\n</divv>\n",
		"<p><div>\n  Something here\n</divv></p>\n",

		"Paragraph\n<div>\nHere? >&<\n</div>\n",
		"<p>Paragraph\n<div>\nHere? &gt;&amp;&lt;\n</div></p>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\n",
		"<p>Paragraph</p>\n\n<div>\nHow about here? >&<\n</div>\n",

		"Paragraph\n<div>\nHere? >&<\n</div>\nAnd here?\n",
		"<p>Paragraph\n<div>\nHere? &gt;&amp;&lt;\n</div>\nAnd here?</p>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\nAnd here?\n",
		"<p>Paragraph</p>\n\n<p><div>\nHow about here? &gt;&amp;&lt;\n</div>\nAnd here?</p>\n",

		"Paragraph\n<div>\nHere? >&<\n</div>\n\nAnd here?\n",
		"<p>Paragraph\n<div>\nHere? &gt;&amp;&lt;\n</div></p>\n\n<p>And here?</p>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\n\nAnd here?\n",
		"<p>Paragraph</p>\n\n<div>\nHow about here? >&<\n</div>\n\n<p>And here?</p>\n",
	}
	doTestsBlock(t, tests, 0)
}

func TestPreformattedHtmlLax(t *testing.T) {
	var tests = []string{
		"Paragraph\n<div>\nHere? >&<\n</div>\n",
		"<p>Paragraph</p>\n\n<div>\nHere? >&<\n</div>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\n",
		"<p>Paragraph</p>\n\n<div>\nHow about here? >&<\n</div>\n",

		"Paragraph\n<div>\nHere? >&<\n</div>\nAnd here?\n",
		"<p>Paragraph</p>\n\n<div>\nHere? >&<\n</div>\n\n<p>And here?</p>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\nAnd here?\n",
		"<p>Paragraph</p>\n\n<div>\nHow about here? >&<\n</div>\n\n<p>And here?</p>\n",

		"Paragraph\n<div>\nHere? >&<\n</div>\n\nAnd here?\n",
		"<p>Paragraph</p>\n\n<div>\nHere? >&<\n</div>\n\n<p>And here?</p>\n",

		"Paragraph\n\n<div>\nHow about here? >&<\n</div>\n\nAnd here?\n",
		"<p>Paragraph</p>\n\n<div>\nHow about here? >&<\n</div>\n\n<p>And here?</p>\n",
	}
	doTestsBlock(t, tests, EXTENSION_LAX_HTML_BLOCKS)
}

func TestFencedCodeBlock(t *testing.T) {
	var tests = []string{
		"``` go\nfunc foo() bool {\n\treturn true;\n}\n```\n",
		"<pre><code class=\"go\">func foo() bool {\n    return true;\n}\n</code></pre>\n",

		"``` c\n/* special & char < > \" escaping */\n```\n",
		"<pre><code class=\"c\">/* special &amp; char &lt; &gt; &quot; escaping */\n</code></pre>\n",

		"``` c\nno *inline* processing ~~of text~~\n```\n",
		"<pre><code class=\"c\">no *inline* processing ~~of text~~\n</code></pre>\n",

		"```\nNo language\n```\n",
		"<pre><code>No language\n</code></pre>\n",

		"``` {ocaml}\nlanguage in braces\n```\n",
		"<pre><code class=\"ocaml\">language in braces\n</code></pre>\n",

		"```    {ocaml}      \nwith extra whitespace\n```\n",
		"<pre><code class=\"ocaml\">with extra whitespace\n</code></pre>\n",

		"```{   ocaml   }\nwith extra whitespace\n```\n",
		"<pre><code class=\"ocaml\">with extra whitespace\n</code></pre>\n",

		"~ ~~ java\nWith whitespace\n~~~\n",
		"<p>~ ~~ java\nWith whitespace\n~~~</p>\n",

		"~~\nonly two\n~~\n",
		"<p>~~\nonly two\n~~</p>\n",

		"```` python\nextra\n````\n",
		"<pre><code class=\"python\">extra\n</code></pre>\n",

		"~~~ perl\nthree to start, four to end\n~~~~\n",
		"<p>~~~ perl\nthree to start, four to end\n~~~~</p>\n",

		"~~~~ perl\nfour to start, three to end\n~~~\n",
		"<p>~~~~ perl\nfour to start, three to end\n~~~</p>\n",

		"~~~ bash\ntildes\n~~~\n",
		"<pre><code class=\"bash\">tildes\n</code></pre>\n",

		"``` lisp\nno ending\n",
		"<p>``` lisp\nno ending</p>\n",

		"~~~ lisp\nend with language\n~~~ lisp\n",
		"<p>~~~ lisp\nend with language\n~~~ lisp</p>\n",

		"```\nmismatched begin and end\n~~~\n",
		"<p>```\nmismatched begin and end\n~~~</p>\n",

		"~~~\nmismatched begin and end\n```\n",
		"<p>~~~\nmismatched begin and end\n```</p>\n",

		"   ``` oz\nleading spaces\n```\n",
		"<pre><code class=\"oz\">leading spaces\n</code></pre>\n",

		"  ``` oz\nleading spaces\n ```\n",
		"<pre><code class=\"oz\">leading spaces\n</code></pre>\n",

		" ``` oz\nleading spaces\n  ```\n",
		"<pre><code class=\"oz\">leading spaces\n</code></pre>\n",

		"``` oz\nleading spaces\n   ```\n",
		"<pre><code class=\"oz\">leading spaces\n</code></pre>\n",
	}

	var tests1 = append(tests, []string{
		"    ``` oz\nleading spaces\n    ```\n",
		"<pre><code>``` oz\n</code></pre>\n\n<p>leading spaces\n    ```</p>\n",
	}...)

	doTestsBlock(t, tests1, EXTENSION_FENCED_CODE)

	var tests2 = append(tests, []string{
		"    ``` oz\nleading spaces\n    ```\n",
		"<pre><code>``` oz\n</code></pre>\n\n<p>leading spaces</p>\n\n<pre><code>```\n</code></pre>\n",
	}...)

	doTestsBlock(t, tests2, EXTENSION_FENCED_CODE|EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK)
}

func TestTable(t *testing.T) {
	var tests = []string{
		"a | b\n---|---\nc | d\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td>c</td>\n<td>d</td>\n</tr>\n</tbody>\n</table>\n",

		"a | b\n---|--\nc | d\n",
		"<p>a | b\n---|--\nc | d</p>\n",

		"|a|b|c|d|\n|----|----|----|---|\n|e|f|g|h|\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n<td>c</td>\n<td>d</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td>e</td>\n<td>f</td>\n<td>g</td>\n<td>h</td>\n</tr>\n</tbody>\n</table>\n",

		"*a*|__b__|[c](C)|d\n---|---|---|---\ne|f|g|h\n",
		"<table>\n<thead>\n<tr>\n<td><em>a</em></td>\n<td><strong>b</strong></td>\n<td><a href=\"C\">c</a></td>\n<td>d</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td>e</td>\n<td>f</td>\n<td>g</td>\n<td>h</td>\n</tr>\n</tbody>\n</table>\n",

		"a|b|c\n---|---|---\nd|e|f\ng|h\ni|j|k|l|m\nn|o|p\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n<td>c</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td>d</td>\n<td>e</td>\n<td>f</td>\n</tr>\n\n" +
			"<tr>\n<td>g</td>\n<td>h</td>\n<td></td>\n</tr>\n\n" +
			"<tr>\n<td>i</td>\n<td>j</td>\n<td>k</td>\n</tr>\n\n" +
			"<tr>\n<td>n</td>\n<td>o</td>\n<td>p</td>\n</tr>\n</tbody>\n</table>\n",

		"a|b|c\n---|---|---\n*d*|__e__|f\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n<td>c</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td><em>d</em></td>\n<td><strong>e</strong></td>\n<td>f</td>\n</tr>\n</tbody>\n</table>\n",

		"a|b|c|d\n:--|--:|:-:|---\ne|f|g|h\n",
		"<table>\n<thead>\n<tr>\n<td align=\"left\">a</td>\n<td align=\"right\">b</td>\n" +
			"<td align=\"center\">c</td>\n<td>d</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td align=\"left\">e</td>\n<td align=\"right\">f</td>\n" +
			"<td align=\"center\">g</td>\n<td>h</td>\n</tr>\n</tbody>\n</table>\n",

		"a|b|c\n---|---|---\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n<td>c</td>\n</tr>\n</thead>\n\n<tbody>\n</tbody>\n</table>\n",

		"a| b|c | d | e\n---|---|---|---|---\nf| g|h | i |j\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b</td>\n<td>c</td>\n<td>d</td>\n<td>e</td>\n</tr>\n</thead>\n\n" +
			"<tbody>\n<tr>\n<td>f</td>\n<td>g</td>\n<td>h</td>\n<td>i</td>\n<td>j</td>\n</tr>\n</tbody>\n</table>\n",

		"a|b\\|c|d\n---|---|---\nf|g\\|h|i\n",
		"<table>\n<thead>\n<tr>\n<td>a</td>\n<td>b|c</td>\n<td>d</td>\n</tr>\n</thead>\n\n<tbody>\n<tr>\n<td>f</td>\n<td>g|h</td>\n<td>i</td>\n</tr>\n</tbody>\n</table>\n",
	}
	doTestsBlock(t, tests, EXTENSION_TABLES)
}
