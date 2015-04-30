// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package example2_test provides a unit test to download an RSS
// feed file and validate it worked. This time is mocks the server.
package example2_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	ex2 "github.com/ArdanStudios/gotraining/09-testing/01-testing/example2"
)

const succeed = "\u2713"
const failed = "\u2717"

// feed is mocking the XML document we except to receive.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:atom="http://www.w3.org/2005/Atom"
     xmlns:openSearch="http://a9.com/-/spec/opensearchrss/1.0/" 
     xmlns:blogger="http://schemas.google.com/blogger/2008"
     xmlns:georss="http://www.georss.org/georss"
     xmlns:gd="http://schemas.google.com/g/2005"
     xmlns:thr="http://purl.org/syndication/thread/1.0" version="2.0">
<channel>
    <atom:id>tag:blogger.com,1999:blog-6298089572270107210</atom:id>
    <lastBuildDate>Mon, 27 Apr 2015 12:42:23 +0000</lastBuildDate>
    <title>Going Go Programming</title>
    <description>Golang For The Common Man : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <managingEditor>noreply@blogger.com (William Kennedy)</managingEditor>
    <generator>Blogger</generator>
    <openSearch:totalResults>67</openSearch:totalResults>
    <openSearch:startIndex>1</openSearch:startIndex>
    <openSearch:itemsPerPage>25</openSearch:itemsPerPage>
    <item>
        <guid isPermaLink="false">tag:blogger.com,1999:blog-6298089572270107210.post-3707264680032475251</guid>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <atom:updated>2015-03-15T11:21:09.934-04:00</atom:updated>
        <title>Object Oriented Programming Mechanics</title>
        <description>&lt;div class="p1"&gt;Go is an object oriented programming language. It may not have inheritance, but in this 20 minute video from the Bangalore meetup, I will show you how object oriented programming practices and techniques can be applied to your Go programs. From an object oriented standpoint, Go does provides the ability to add behavior to your types via methods, allows you to implement polymorphic behavior via interfaces and gives you a way to extend the state and behavior of any existing type via type embedding. Go also provides a form of encapsulation that allows your types, including their fields and methods, to be visible or invisible. Everything you need to write object oriented programs is available in Go.&lt;/div&gt;&lt;br /&gt;&lt;iframe allowfullscreen="" frameborder="0" height="510" src="https://www.youtube.com/embed/gRpUfjTwSOo" width="100%"&gt;&lt;/iframe&gt; &lt;br /&gt;&lt;b&gt;&lt;br /&gt;&lt;/b&gt;&lt;b&gt;Example 1 - Methods: &lt;/b&gt;&lt;br /&gt;&lt;a href="http://play.golang.org/p/hfRkC6nKag" target="_blank"&gt;http://play.golang.org/p/hfRkC6nKag&lt;/a&gt;&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Example 2 - Interfaces:&lt;/b&gt;&lt;br /&gt;&lt;a href="http://play.golang.org/p/F1UyKlTh3k" target="_blank"&gt;http://play.golang.org/p/F1UyKlTh3k&lt;/a&gt;&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Example 3 - Extending Types:&lt;/b&gt;&lt;br /&gt;&lt;a href="http://play.golang.org/p/JJ811lBwoz" target="_blank"&gt;http://play.golang.org/p/JJ811lBwoz&lt;/a&gt;&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Example 4 - Overriding Inner Types:&lt;/b&gt;&lt;br /&gt;&lt;a href="http://play.golang.org/p/-xQFBv9_82" target="_blank"&gt;http://play.golang.org/p/-xQFBv9_82&lt;/a&gt;</description>
        <link>http://www.goinggo.net/2015/03/object-oriented-programming-mechanics.html</link>
        <author>noreply@blogger.com (William Kennedy)</author>
        <media:thumbnail xmlns:media="http://search.yahoo.com/mrss/" url="http://img.youtube.com/vi/gRpUfjTwSOo/default.jpg" height="72" width="72"/>
        <thr:total>0</thr:total>
    </item>
</channel>
</rss>`

// mockServer returns a pointer to a server to handle the mock get call.
func mockServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}))
	return server
}

// TestDownload tests if download web content is working.
func TestDownload(t *testing.T) {
	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		resp, err := http.Get(server.URL)
		if err == nil {
			t.Log("\tShould be able to make the Get call.",
				succeed)
		} else {
			t.Fatal("\tShould be able to make the Get call.",
				failed, err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			t.Log("\tShould receive a \"200\" status code.",
				succeed)
		} else {
			t.Error("\tShould receive a \"200\" status code.",
				failed, resp.StatusCode)
		}

		var d ex2.Document
		if err := xml.NewDecoder(resp.Body).Decode(&d); err == nil {
			t.Log("\tShould be able to unmarshal the response.",
				succeed)
		} else {
			t.Fatal("\tShould be able to unmarshal the response.",
				failed, err)
		}

		if len(d.Channel.Items) == 1 {
			t.Log("\tShould have \"1\" item in the feed.",
				succeed)
		} else {
			t.Fatal("\tShould have \"1\" item in the feed.",
				failed, len(d.Channel.Items))
		}
	}
}
