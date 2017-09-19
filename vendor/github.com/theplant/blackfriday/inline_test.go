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

import (
	"testing"
)

func runMarkdownInline(input string) string {
	extensions := 0
	extensions |= EXTENSION_AUTOLINK
	extensions |= EXTENSION_STRIKETHROUGH

	htmlFlags := 0
	htmlFlags |= HTML_USE_XHTML

	renderer := HtmlRenderer(htmlFlags, "", "")

	return string(Markdown([]byte(input), renderer, extensions))
}

func doTestsInline(t *testing.T, tests []string) {
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
		actual := runMarkdownInline(candidate)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}

		// now test every substring to stress test bounds checking
		if !testing.Short() {
			for start := 0; start < len(input); start++ {
				for end := start + 1; end <= len(input); end++ {
					candidate = input[start:end]
					_ = runMarkdownInline(candidate)
				}
			}
		}
	}
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

		"*improper **nesting* is** bad\n",
		"<p><em>improper **nesting</em> is** bad</p>\n",
	}
	doTestsInline(t, tests)
}

func TestStrikeThrough(t *testing.T) {
	var tests = []string{
		"nothing inline\n",
		"<p>nothing inline</p>\n",

		"simple ~~inline~~ test\n",
		"<p>simple <del>inline</del> test</p>\n",

		"~~at the~~ beginning\n",
		"<p><del>at the</del> beginning</p>\n",

		"at the ~~end~~\n",
		"<p>at the <del>end</del></p>\n",

		"~~try two~~ in ~~one line~~\n",
		"<p><del>try two</del> in <del>one line</del></p>\n",

		"over ~~two\nlines~~ test\n",
		"<p>over <del>two\nlines</del> test</p>\n",

		"odd ~~number of~~ markers~~ here\n",
		"<p>odd <del>number of</del> markers~~ here</p>\n",

		"odd ~~number\nof~~ markers~~ here\n",
		"<p>odd <del>number\nof</del> markers~~ here</p>\n",
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
