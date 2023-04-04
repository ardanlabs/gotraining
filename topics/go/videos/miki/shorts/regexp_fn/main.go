// regexp can use functions for substitution.

package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\$[A-Z_]+`)
	urlTemplate := `https://$HOST:$PORT`

	conf := map[string]string{
		"HOST": "www.ardanlabs.com",
		"PORT": "443",
	}

	sub := func(match string) string {
		key := match[1:] // Remove $ prefix
		return conf[key]
	}

	url := re.ReplaceAllStringFunc(urlTemplate, sub)
	fmt.Println(url)
	// https://www.ardanlabs.com:443
}
