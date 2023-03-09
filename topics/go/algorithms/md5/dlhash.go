package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func urlSig(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%q: bad status - %s", url, resp.Status)
	}

	// Find MD5 hash in HTTP headers.
	const (
		header = "x-goog-hash"
		prefix = "md5="
	)
	b64hash := ""
	values := resp.Header.Values(header)
	for _, v := range values {
		if strings.HasPrefix(v, prefix) {
			b64hash = v[len(prefix):]
			break
		}
	}

	if b64hash == "" {
		return "", fmt.Errorf("can't find md5 hash %s: %v", header, values)
	}

	hash, err := base64.StdEncoding.DecodeString(b64hash)
	if err != nil {
		return "", err
	}

	// Convert hash to "eec1fa5ce8077d7030e194eb5989c937" format.
	return fmt.Sprintf("%x", hash), nil
}

func main() {
	url := "https://storage.googleapis.com/gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_B2.TIF"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(urlSig(ctx, url))
}
