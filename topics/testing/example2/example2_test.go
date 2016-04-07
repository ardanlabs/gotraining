// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a basic unit table test.
package example2

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload validates the http Get function can download content and
// handles different status conditions properly.
func TestDownload(t *testing.T) {
	urls := []struct {
		url        string
		statusCode int
	}{
		{"http://www.goinggo.net/feeds/posts/default?alt=rss", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, u := range urls {
			t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
				}
				t.Logf("\t%s\tShould be able to make the Get call.", succeed)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t%s\tShould receive a %d status code.", succeed, u.statusCode)
				} else {
					t.Errorf("\t%s\tShould receive a %d status code : %v", failed, u.statusCode, resp.StatusCode)
				}
			}
		}
	}
}
