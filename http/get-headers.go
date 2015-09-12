package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var address string

func init() {
	const (
		defaultAddress = "http://golang.org"
		usage          = "the URL or IP address"
	)
	flag.StringVar(&address, "address", defaultAddress, usage)
	flag.StringVar(&address, "a", defaultAddress, usage+" (shorthand)")
}

func main() {
	flag.Parse()

	g, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", g.Header)
}
