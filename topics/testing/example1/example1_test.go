// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a basic unit test.
package example1

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload validates the http Get function can download content.
func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					failed, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status code. %v",
					statusCode, succeed)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status code. %v %v",
					statusCode, failed, resp.StatusCode)
			}
		}
	}
}
