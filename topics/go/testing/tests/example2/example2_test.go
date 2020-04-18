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
	tt := []struct {
		url        string
		statusCode int
	}{
		{"https://www.ardanlabs.com/blog/index.xml", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for testID, test := range tt {
			t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, test.url, test.statusCode)
			{
				resp, err := http.Get(test.url)
				if err != nil {
					t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call : %v", failed, testID, err)
				}
				t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, testID)

				defer resp.Body.Close()

				if resp.StatusCode == test.statusCode {
					t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, testID, test.statusCode)
				} else {
					t.Errorf("\t%s\tTest %d:\tShould receive a %d status code : %v", failed, testID, test.statusCode, resp.StatusCode)
				}
			}
		}
	}
}
