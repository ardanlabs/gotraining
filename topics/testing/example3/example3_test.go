// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to mock an HTTP GET call internally.
package example3

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// feed is mocking the XML document we expect to receive.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// Item defines the fields associated with the item tag in
// the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// mockServer returns a pointer to a server to handle the mock get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload validates the http Get function can download content and
// the content can be unmarshaled and clean.
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\tShould receive a %d status code : %v", failed, statusCode, resp.StatusCode)
			}
			t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the response.", succeed)

			if len(d.Channel.Items) == 1 {
				t.Logf("\t%s\tShould have 1 item in the feed.", succeed)
			} else {
				t.Errorf("\t%s\tShould have 1 item in the feed : %d", failed, len(d.Channel.Items))
			}
		}
	}
}
