package main

import (
	"fmt"
	urllib "net/url"
)

func validateURL(url string) error {
	_, err := urllib.Parse(url)
	return err
}

func main() {
	fmt.Println(validateURL("https://go.dev"))
	fmt.Println(validateURL("://"))

}
