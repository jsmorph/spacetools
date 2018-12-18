// Package main is simple tool for getting geomagnetic data in JSON.
//
// See https://www.swpc.noaa.gov/products/planetary-k-index,
// https://services.swpc.noaa.gov/text/daily-geomagnetic-indices.txt,
// and ftp://ftp.swpc.noaa.gov/pub/indices/old_indices/.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/jsonschema"
)

func main() {

	var (
		filename = flag.String("f", "", "filename (if data is local)")
		url      = flag.String("u", "https://services.swpc.noaa.gov/text/daily-geomagnetic-indices.txt", "URL for file")
		debug    = flag.Bool("d", false, "include lines with records")
		array    = flag.Bool("a", false, "wrap records in an array")
		schema   = flag.Bool("s", false, "just print out JSON Schema and stop")
		help     = flag.Bool("h", false, "get help")
	)

	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Get daily geomagnetic data in JSON\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		return
	}

	if *schema {
		var x interface{}
		if *array {
			x = []Day{}
		} else {
			x = Day{}
		}
		s := jsonschema.Reflect(x)
		js, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", js)
		return
	}

	var bs []byte
	var err error

	if len(*filename) == 0 {
		var resp *http.Response
		log.Printf("downloading %s", *url)
		if resp, err = http.Get(*url); err == nil {
			if bs, err = ioutil.ReadAll(resp.Body); err == nil {
				err = ioutil.WriteFile("data.txt", bs, 0644)
			}
		}
	} else {

		switch *filename {
		case "-":
			bs, err = ioutil.ReadAll(os.Stdin)
		default:

			bs, err = ioutil.ReadFile(*filename)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	ds, err := ParseFile(string(bs))
	if err != nil {
		log.Fatal(err)
	}

	if !*debug {
		for _, d := range ds {
			d.Line = ""
		}
	}

	pad := ""
	if *array {
		fmt.Printf("[\n")
		pad = "  "
	}
	last := len(ds) - 1
	for i, d := range ds {
		js, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s%s", pad, js)
		if *array && i < last {
			fmt.Printf(",")
		}
		fmt.Printf("\n")
	}
	if *array {
		fmt.Printf("]\n")
	}
}
