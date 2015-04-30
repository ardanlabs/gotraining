// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

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
		t.Logf("\tWhen the URL is \"%s\" with status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err == nil {
				t.Log("\t\tShould be able to make the Get call.",
					succeed)
			} else {
				t.Fatal("\t\tShould be able to make the Get call.",
					failed, err)
			}

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
