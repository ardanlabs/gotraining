//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
// Unit tests for inline parsing
//

package blackfriday

import "testing"

func runMarkdownInline(input string, extensions, htmlFlags int) string {
	extensions |= EXTENSION_AUTOLINK
	extensions |= EXTENSION_STRIKETHROUGH

	htmlFlags |= HTML_USE_XHTML

	renderer := HtmlRenderer(htmlFlags, "", "")

	return string(Markdown([]byte(input), renderer, extensions))
}

func doTestsInline(t *testing.T, tests []string) {
	doTestsInlineParam(t, tests, 0, 0)
}

func doTestsInlineParam(t *testing.T, tests []string, extensions, htmlFlags int) {
	// catch and report panics
	var candidate string
	/*
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("\npanic while processing [%#v] (%v)\n", candidate, err)
			}
		}()
	*/

	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		candidate = input
		expected := tests[i+1]
		actual := runMarkdownInline(candidate, extensions, htmlFlags)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}

		// now test every substring to stress test bounds checking
		if !testing.Short() {
			for start := 0; start < len(input); start++ {
				for end := start + 1; end <= len(input); end++ {
					candidate = input[start:end]
					_ = runMarkdownInline(candidate, extensions, htmlFlags)
				}
			}
		}
	}
}

func TestRawHtmlTag(t *testing.T) {
	tests := []string{
		"zz <style>p {}</style>\n",
		"<p>zz p {}</p>\n",

		"zz <STYLE>p {}</STYLE>\n",
		"<p>zz p {}</p>\n",

		"<SCRIPT>alert()</SCRIPT>\n",
		"<p>alert()</p>\n",

		"zz <SCRIPT>alert()</SCRIPT>\n",
		"<p>zz alert()</p>\n",

		"zz <script>alert()</script>\n",
		"<p>zz alert()</p>\n",

		" <script>alert()</script>\n",
		"<p>alert()</p>\n",

		"<script>alert()</script>\n",
		"<p>alert()</p>\n",

		"<script src='foo'></script>\n",
		"<p></p>\n",

		"zz <script src='foo'></script>\n",
		"<p>zz </p>\n",

		"zz <script src=foo></script>\n",
		"<p>zz </p>\n",
	}
	doTestsInlineParam(t, tests, 0, HTML_SKIP_STYLE|HTML_SKIP_SCRIPT)
}

func TestEmphasis(t *testing.T) {
	var tests = []string{
		"nothing inline\n",
		"<p>nothing inline</p>\n",

		"simple *inline* test\n",
		"<p>simple <em>inline</em> test</p>\n",

		"*at the* beginning\n",
		"<p><em>at the</em> beginning</p>\n",

		"at the *end*\n",
		"<p>at the <em>end</em></p>\n",

		"*try two* in *one line*\n",
		"<p><em>try two</em> in <em>one line</em></p>\n",

		"over *two\nlines* test\n",
		"<p>over <em>two\nlines</em> test</p>\n",

		"odd *number of* markers* here\n",
		"<p>odd <em>number of</em> markers* here</p>\n",

		"odd *number\nof* markers* here\n",
		"<p>odd <em>number\nof</em> markers* here</p>\n",

		"marker is *next to character*at the right\n",
		"<p>marker is <em>next to character</em>at the right</p>\n",

		"marker is*next to character* at the left\n",
		"<p>marker is<em>next to character</em> at the left</p>\n",

		"marker is*next to character*at both side\n",
		"<p>marker is<em>next to character</em>at both side</p>\n",

		"simple _inline_ test\n",
		"<p>simple <em>inline</em> test</p>\n",

		"_at the_ beginning\n",
		"<p><em>at the</em> beginning</p>\n",

		"at the _end_\n",
		"<p>at the <em>end</em></p>\n",

		"_try two_ in _one line_\n",
		"<p><em>try two</em> in <em>one line</em></p>\n",

		"over _two\nlines_ test\n",
		"<p>over <em>two\nlines</em> test</p>\n",

		"odd _number of_ markers_ here\n",
		"<p>odd <em>number of</em> markers_ here</p>\n",

		"odd _number\nof_ markers_ here\n",
		"<p>odd <em>number\nof</em> markers_ here</p>\n",

		"marker is _next to character_at the right\n",
		"<p>marker is <em>next to character</em>at the right</p>\n",

		"marker is_next to character_ at the left\n",
		"<p>marker is<em>next to character</em> at the left</p>\n",

		"marker is_next to character_at both side\n",
		"<p>marker is<em>next to character</em>at both side</p>\n",

		"mix of *markers_\n",
		"<p>mix of *markers_</p>\n",
	}
	doTestsInline(t, tests)
}

func TestStrong(t *testing.T) {
	var tests = []string{
		"nothing inline\n",
		"<p>nothing inline</p>\n",

		"simple **inline** test\n",
		"<p>simple <strong>inline</strong> test</p>\n",

		"**at the** beginning\n",
		"<p><strong>at the</strong> beginning</p>\n",

		"at the **end**\n",
		"<p>at the <strong>end</strong></p>\n",

		"**try two** in **one line**\n",
		"<p><strong>try two</strong> in <strong>one line</strong></p>\n",

		"over **two\nlines** test\n",
		"<p>over <strong>two\nlines</strong> test</p>\n",

		"odd **number of** markers** here\n",
		"<p>odd <strong>number of</strong> markers** here</p>\n",

		"odd **number\nof** markers** here\n",
		"<p>odd <strong>number\nof</strong> markers** here</p>\n",

		"simple __inline__ test\n",
		"<p>simple <strong>inline</strong> test</p>\n",

		"__at the__ beginning\n",
		"<p><strong>at the</strong> beginning</p>\n",

		"at the __end__\n",
		"<p>at the <strong>end</strong></p>\n",

		"__try two__ in __one line__\n",
		"<p><strong>try two</strong> in <strong>one line</strong></p>\n",

		"over __two\nlines__ test\n",
		"<p>over <strong>two\nlines</strong> test</p>\n",

		"odd __number of__ markers__ here\n",
		"<p>odd <strong>number of</strong> markers__ here</p>\n",

		"odd __number\nof__ markers__ here\n",
		"<p>odd <strong>number\nof</strong> markers__ here</p>\n",

		"mix of **markers__\n",
		"<p>mix of **markers__</p>\n",
	}
	doTestsInline(t, tests)
}

func TestEmphasisMix(t *testing.T) {
	var tests = []string{
		"***triple emphasis***\n",
		"<p><strong><em>triple emphasis</em></strong></p>\n",

		"***triple\nemphasis***\n",
		"<p><strong><em>triple\nemphasis</em></strong></p>\n",

		"___triple emphasis___\n",
		"<p><strong><em>triple emphasis</em></strong></p>\n",

		"***triple emphasis___\n",
		"<p>***triple emphasis___</p>\n",

		"*__triple emphasis__*\n",
		"<p><em><strong>triple emphasis</strong></em></p>\n",

		"__*triple emphasis*__\n",
		"<p><strong><em>triple emphasis</em></strong></p>\n",

		"**improper *nesting** is* bad\n",
		"<p><strong>improper *nesting</strong> is* bad</p>\n",

		"***improper *nesting*** is* bad\n",
		"<p><strong><em>improper *nesting</em></strong> is* bad</p>\n",

		"*improper **nesting* is** bad\n",
		"<p><em>improper </em><em>nesting</em> is** bad</p>\n",
	}
	doTestsInline(t, tests)
}

func TestStrikeThrough(t *testing.T) {
	var tests = []string{
		"nothing inline\n",
		"<p>nothing inline</p>\n",

		"simple ~~inline~~ test\n",
		"<p>simple <s>inline</s> test</p>\n",

		"~~at the~~ beginning\n",
		"<p><s>at the</s> beginning</p>\n",

		"at the ~~end~~\n",
		"<p>at the <s>end</s></p>\n",

		"~~try two~~ in ~~one line~~\n",
		"<p><s>try two</s> in <s>one line</s></p>\n",

		"over ~~two\nlines~~ test\n",
		"<p>over <s>two\nlines</s> test</p>\n",

		"odd ~~number of~~ markers~~ here\n",
		"<p>odd <s>number of</s> markers~~ here</p>\n",

		"odd ~~number\nof~~ markers~~ here\n",
		"<p>odd <s>number\nof</s> markers~~ here</p>\n",
	}
	doTestsInline(t, tests)
}

func TestCodeSpan(t *testing.T) {
	var tests = []string{
		"`source code`\n",
		"<p><code>source code</code></p>\n",

		"` source code with spaces `\n",
		"<p><code>source code with spaces</code></p>\n",

		"` source code with spaces `not here\n",
		"<p><code>source code with spaces</code>not here</p>\n",

		"a `single marker\n",
		"<p>a `single marker</p>\n",

		"a single multi-tick marker with ``` no text\n",
		"<p>a single multi-tick marker with ``` no text</p>\n",

		"markers with ` ` a space\n",
		"<p>markers with  a space</p>\n",

		"`source code` and a `stray\n",
		"<p><code>source code</code> and a `stray</p>\n",

		"`source *with* _awkward characters_ in it`\n",
		"<p><code>source *with* _awkward characters_ in it</code></p>\n",

		"`split over\ntwo lines`\n",
		"<p><code>split over\ntwo lines</code></p>\n",

		"```multiple ticks``` for the marker\n",
		"<p><code>multiple ticks</code> for the marker</p>\n",

		"```multiple ticks `with` ticks inside```\n",
		"<p><code>multiple ticks `with` ticks inside</code></p>\n",
	}
	doTestsInline(t, tests)
}

func TestLineBreak(t *testing.T) {
	var tests = []string{
		"this line  \nhas a break\n",
		"<p>this line<br />\nhas a break</p>\n",

		"this line \ndoes not\n",
		"<p>this line\ndoes not</p>\n",

		"this has an   \nextra space\n",
		"<p>this has an<br />\nextra space</p>\n",
	}
	doTestsInline(t, tests)
}

func TestInlineLink(t *testing.T) {
	var tests = []string{
		"[foo](/bar/)\n",
		"<p><a href=\"/bar/\">foo</a></p>\n",

		"[foo with a title](/bar/ \"title\")\n",
		"<p><a href=\"/bar/\" title=\"title\">foo with a title</a></p>\n",

		"[foo with a title](/bar/\t\"title\")\n",
		"<p><a href=\"/bar/\" title=\"title\">foo with a title</a></p>\n",

		"[foo with a title](/bar/ \"title\"  )\n",
		"<p><a href=\"/bar/\" title=\"title\">foo with a title</a></p>\n",

		"[foo with a title](/bar/ title with no quotes)\n",
		"<p><a href=\"/bar/ title with no quotes\">foo with a title</a></p>\n",

		"[foo]()\n",
		"<p>[foo]()</p>\n",

		"![foo](/bar/)\n",
		"<p><img src=\"/bar/\" alt=\"foo\" />\n</p>\n",

		"![foo with a title](/bar/ \"title\")\n",
		"<p><img src=\"/bar/\" alt=\"foo with a title\" title=\"title\" />\n</p>\n",

		"![foo with a title](/bar/\t\"title\")\n",
		"<p><img src=\"/bar/\" alt=\"foo with a title\" title=\"title\" />\n</p>\n",

		"![foo with a title](/bar/ \"title\"  )\n",
		"<p><img src=\"/bar/\" alt=\"foo with a title\" title=\"title\" />\n</p>\n",

		"![foo with a title](/bar/ title with no quotes)\n",
		"<p><img src=\"/bar/ title with no quotes\" alt=\"foo with a title\" />\n</p>\n",

		"![foo]()\n",
		"<p>![foo]()</p>\n",

		"[a link]\t(/with_a_tab/)\n",
		"<p><a href=\"/with_a_tab/\">a link</a></p>\n",

		"[a link]  (/with_spaces/)\n",
		"<p><a href=\"/with_spaces/\">a link</a></p>\n",

		"[text (with) [[nested] (brackets)]](/url/)\n",
		"<p><a href=\"/url/\">text (with) [[nested] (brackets)]</a></p>\n",

		"[text (with) [broken nested] (brackets)]](/url/)\n",
		"<p>[text (with) <a href=\"brackets\">broken nested</a>]](/url/)</p>\n",

		"[text\nwith a newline](/link/)\n",
		"<p><a href=\"/link/\">text\nwith a newline</a></p>\n",

		"[text in brackets] [followed](/by a link/)\n",
		"<p>[text in brackets] <a href=\"/by a link/\">followed</a></p>\n",

		"[link with\\] a closing bracket](/url/)\n",
		"<p><a href=\"/url/\">link with] a closing bracket</a></p>\n",

		"[link with\\[ an opening bracket](/url/)\n",
		"<p><a href=\"/url/\">link with[ an opening bracket</a></p>\n",

		"[link with\\) a closing paren](/url/)\n",
		"<p><a href=\"/url/\">link with) a closing paren</a></p>\n",

		"[link with\\( an opening paren](/url/)\n",
		"<p><a href=\"/url/\">link with( an opening paren</a></p>\n",

		"[link](  with whitespace)\n",
		"<p><a href=\"with whitespace\">link</a></p>\n",

		"[link](  with whitespace   )\n",
		"<p><a href=\"with whitespace\">link</a></p>\n",

		"[![image](someimage)](with image)\n",
		"<p><a href=\"with image\"><img src=\"someimage\" alt=\"image\" />\n</a></p>\n",

		"[link](url \"one quote)\n",
		"<p><a href=\"url &quot;one quote\">link</a></p>\n",

		"[link](url 'one quote)\n",
		"<p><a href=\"url 'one quote\">link</a></p>\n",

		"[link](<url>)\n",
		"<p><a href=\"url\">link</a></p>\n",

		"[link & ampersand](/url/)\n",
		"<p><a href=\"/url/\">link &amp; ampersand</a></p>\n",

		"[link &amp; ampersand](/url/)\n",
		"<p><a href=\"/url/\">link &amp; ampersand</a></p>\n",

		"[link](/url/&query)\n",
		"<p><a href=\"/url/&amp;query\">link</a></p>\n",

		"[[t]](/t)\n",
		"<p><a href=\"/t\">[t]</a></p>\n",
	}
	doTestsInline(t, tests)
}

func TestReferenceLink(t *testing.T) {
	var tests = []string{
		"[link][ref]\n",
		"<p>[link][ref]</p>\n",

		"[link][ref]\n   [ref]: /url/ \"title\"\n",
		"<p><a href=\"/url/\" title=\"title\">link</a></p>\n",

		"[link][ref]\n   [ref]: /url/\n",
		"<p><a href=\"/url/\">link</a></p>\n",

		"   [ref]: /url/\n",
		"",

		"   [ref]: /url/\n[ref2]: /url/\n [ref3]: /url/\n",
		"",

		"   [ref]: /url/\n[ref2]: /url/\n [ref3]: /url/\n    [4spaces]: /url/\n",
		"<pre><code>[4spaces]: /url/\n</code></pre>\n",

		"[hmm](ref2)\n   [ref]: /url/\n[ref2]: /url/\n [ref3]: /url/\n",
		"<p><a href=\"ref2\">hmm</a></p>\n",

		"[ref]\n",
		"<p>[ref]</p>\n",

		"[ref]\n   [ref]: /url/ \"title\"\n",
		"<p><a href=\"/url/\" title=\"title\">ref</a></p>\n",
	}
	doTestsInline(t, tests)
}

func TestItalicAndBoldAreNeighbor(t *testing.T) {
	var tests = []string{
		//"*Italic* **Bold**\n",
		//"<p><em>Italic</em> <strong>Bold</strong></p>\n",
		"*Italic** hey\n",
		"<p><em>Italic</em>* hey</p>\n",

		"*Italic***Bold**\n",
		"<p><em>Italic</em><strong>Bold</strong></p>\n",
	}
	doTestsInline(t, tests)
}

func TestTags(t *testing.T) {
	var tests = []string{
		"a <span>tag</span>\n",
		"<p>a <span>tag</span></p>\n",

		"<span>tag</span>\n",
		"<p><span>tag</span></p>\n",

		"<span>mismatch</spandex>\n",
		"<p><span>mismatch</spandex></p>\n",

		"a <singleton /> tag\n",
		"<p>a <singleton /> tag</p>\n",
	}
	doTestsInline(t, tests)
}

func TestAutoLink(t *testing.T) {
	var tests = []string{
		"http://foo.com/\n",
		"<p><a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"1 http://foo.com/\n",
		"<p>1 <a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"1http://foo.com/\n",
		"<p>1<a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"1.http://foo.com/\n",
		"<p>1.<a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"1. http://foo.com/\n",
		"<ol>\n<li><a href=\"http://foo.com/\">http://foo.com/</a></li>\n</ol>\n",

		"-http://foo.com/\n",
		"<p>-<a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"- http://foo.com/\n",
		"<ul>\n<li><a href=\"http://foo.com/\">http://foo.com/</a></li>\n</ul>\n",

		"_http://foo.com/\n",
		"<p>_<a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"令狐http://foo.com/\n",
		"<p>令狐<a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"令狐 http://foo.com/\n",
		"<p>令狐 <a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"ahttp://foo.com/\n",
		"<p>ahttp://foo.com/</p>\n",

		">http://foo.com/\n",
		"<blockquote>\n<p><a href=\"http://foo.com/\">http://foo.com/</a></p>\n</blockquote>\n",

		"> http://foo.com/\n",
		"<blockquote>\n<p><a href=\"http://foo.com/\">http://foo.com/</a></p>\n</blockquote>\n",

		"go to <http://foo.com/>\n",
		"<p>go to <a href=\"http://foo.com/\">http://foo.com/</a></p>\n",

		"a secure <https://link.org>\n",
		"<p>a secure <a href=\"https://link.org\">https://link.org</a></p>\n",

		"an email <mailto:some@one.com>\n",
		"<p>an email <a href=\"mailto:some@one.com\">some@one.com</a></p>\n",

		"an email <mailto://some@one.com>\n",
		"<p>an email <a href=\"mailto://some@one.com\">some@one.com</a></p>\n",

		"an email <some@one.com>\n",
		"<p>an email <a href=\"mailto:some@one.com\">some@one.com</a></p>\n",

		"an ftp <ftp://old.com>\n",
		"<p>an ftp <a href=\"ftp://old.com\">ftp://old.com</a></p>\n",

		"an ftp <ftp:old.com>\n",
		"<p>an ftp <a href=\"ftp:old.com\">ftp:old.com</a></p>\n",

		"a link with <http://new.com?query=foo&bar>\n",
		"<p>a link with <a href=\"http://new.com?query=foo&amp;bar\">" +
			"http://new.com?query=foo&amp;bar</a></p>\n",

		"quotes mean a tag <http://new.com?query=\"foo\"&bar>\n",
		"<p>quotes mean a tag <http://new.com?query=\"foo\"&bar></p>\n",

		"quotes mean a tag <http://new.com?query='foo'&bar>\n",
		"<p>quotes mean a tag <http://new.com?query='foo'&bar></p>\n",

		"unless escaped <http://new.com?query=\\\"foo\\\"&bar>\n",
		"<p>unless escaped <a href=\"http://new.com?query=&quot;foo&quot;&amp;bar\">" +
			"http://new.com?query=&quot;foo&quot;&amp;bar</a></p>\n",

		"even a > can be escaped <http://new.com?q=\\>&etc>\n",
		"<p>even a &gt; can be escaped <a href=\"http://new.com?q=&gt;&amp;etc\">" +
			"http://new.com?q=&gt;&amp;etc</a></p>\n",
	}
	doTestsInline(t, tests)
}

func TestFootnotes(t *testing.T) {
	tests := []string{
		"testing footnotes.[^a]\n\n[^a]: This is the note\n",
		`<p>testing footnotes.<sup class="footnote-ref" id="fnref:a"><a rel="footnote" href="#fn:a">1</a></sup></p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:a">This is the note
</li>
</ol>
</div>
`,

		`testing long[^b] notes.

[^b]: Paragraph 1

	Paragraph 2
	
	` + "```\n\tsome code\n\t```" + `
	
	Paragraph 3

No longer in the footnote
`,
		`<p>testing long<sup class="footnote-ref" id="fnref:b"><a rel="footnote" href="#fn:b">1</a></sup> notes.</p>

<p>No longer in the footnote</p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:b"><p>Paragraph 1</p>

<p>Paragraph 2</p>

<p><code>
some code
</code></p>

<p>Paragraph 3</p>
</li>
</ol>
</div>
`,

		`testing[^c] multiple[^d] notes.

[^c]: this is [note] c


omg

[^d]: this is note d

what happens here

[note]: /link/c

`,
		`<p>testing<sup class="footnote-ref" id="fnref:c"><a rel="footnote" href="#fn:c">1</a></sup> multiple<sup class="footnote-ref" id="fnref:d"><a rel="footnote" href="#fn:d">2</a></sup> notes.</p>

<p>omg</p>

<p>what happens here</p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:c">this is <a href="/link/c">note</a> c
</li>
<li id="fn:d">this is note d
</li>
</ol>
</div>
`,

		"testing inline^[this is the note] notes.\n",
		`<p>testing inline<sup class="footnote-ref" id="fnref:this-is-the-note"><a rel="footnote" href="#fn:this-is-the-note">1</a></sup> notes.</p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:this-is-the-note">this is the note</li>
</ol>
</div>
`,

		"testing multiple[^1] types^[inline note] of notes[^2]\n\n[^2]: the second deferred note\n[^1]: the first deferred note\n\n\twhich happens to be a block\n",
		`<p>testing multiple<sup class="footnote-ref" id="fnref:1"><a rel="footnote" href="#fn:1">1</a></sup> types<sup class="footnote-ref" id="fnref:inline-note"><a rel="footnote" href="#fn:inline-note">2</a></sup> of notes<sup class="footnote-ref" id="fnref:2"><a rel="footnote" href="#fn:2">3</a></sup></p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:1"><p>the first deferred note</p>

<p>which happens to be a block</p>
</li>
<li id="fn:inline-note">inline note</li>
<li id="fn:2">the second deferred note
</li>
</ol>
</div>
`,

		`This is a footnote[^1]^[and this is an inline footnote]

[^1]: the footnote text.

    may be multiple paragraphs.
`,
		`<p>This is a footnote<sup class="footnote-ref" id="fnref:1"><a rel="footnote" href="#fn:1">1</a></sup><sup class="footnote-ref" id="fnref:and-this-is-an-i"><a rel="footnote" href="#fn:and-this-is-an-i">2</a></sup></p>
<div class="footnotes">

<hr />

<ol>
<li id="fn:1"><p>the footnote text.</p>

<p>may be multiple paragraphs.</p>
</li>
<li id="fn:and-this-is-an-i">and this is an inline footnote</li>
</ol>
</div>
`,

		"empty footnote[^]\n\n[^]: fn text",
		"<p>empty footnote<sup class=\"footnote-ref\" id=\"fnref:\"><a rel=\"footnote\" href=\"#fn:\">1</a></sup></p>\n<div class=\"footnotes\">\n\n<hr />\n\n<ol>\n<li id=\"fn:\">fn text\n</li>\n</ol>\n</div>\n",
	}

	doTestsInlineParam(t, tests, EXTENSION_FOOTNOTES, 0)
}
