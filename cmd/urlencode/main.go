// urlencode URL-encodes one or more command-line arguments emitting each to STDOUT.
package main

import (
	"flag"
	"net/url"
	"fmt"
)

func main() {

	flag.Parse()

	for _, str := range flag.Args() {
		enc := url.QueryEscape(str)
		fmt.Println(enc)
	}
}
