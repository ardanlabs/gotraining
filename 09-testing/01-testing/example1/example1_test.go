// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package example1_test provides a unit test to download an RSS
// feed file and validate it worked.
package example1

import (
	"encoding/xml"
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload tests if download web content is working.
func TestDownload(t *testing.T) {
	URL := "http://www.goinggo.net/feeds/posts/default?alt=rss"

	t.Log("Given the need to test downloading content.")
	{
		resp, err := http.Get(URL)
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

		var d Document
		if err := xml.NewDecoder(resp.Body).Decode(&d); err == nil {
			t.Log("\tShould be able to unmarshal the response.",
				succeed)
		} else {
			t.Fatal("\tShould be able to unmarshal the response.",
				failed, err)
		}

		if len(d.Channel.Items) == 25 {
			t.Log("\tShould have \"25\" item in the feed.",
				succeed)
		} else {
			t.Fatal("\tShould have \"25\" item in the feed.",
				failed, len(d.Channel.Items))
		}
	}
}
