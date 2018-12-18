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
	"regexp"
	"strings"

	"github.com/alecthomas/jsonschema"
)

// HeaderPattern is a regular expression that should match the header
// of a reasonable input file.
//
// See https://services.swpc.noaa.gov/text/daily-geomagnetic-indices.txt.
var HeaderPattern = `:Product: Daily Geomagnetic Data.*
:Issued: .*
#
#  Prepared by the U.S. Dept. of Commerce, NOAA, Space Weather Prediction Center
#  Please send comment and suggestions to SWPC.Webmaster@noaa.gov
#
#  .*
#
#
#                Middle Latitude        High Latitude            Estimated
#              - Fredericksburg -     ---- College ----      --- Planetary ---
#  Date        A     K-indices        A     K-indices        A     K-indices
`

// Site represents data for one location (e.g., Fredericksburg).
//
// See ftp://ftp.swpc.noaa.gov/pub/indices/old_indices/README:
//
// The daily 24-hour A index and eight 3-hourly K indices from the
// Fredericksburg (middle-latitude) and College  (high-latitude) stations
// monitoring Earth's magnetic field. The estimated planetary 24 hour
// A index and eight 3-hourly K indices are derived in real time from a
// network of western hemisphere ground-based magnetometers.  These indices
// may differ from the final Ap values derived by the Institut fur
// Geophysik, Gottingen, Germany, using a global network of magnetometers.
// K indices range from 0 (very quiet) to 9 (extremely disturbed).  A indices
// range from 0 (very quiet) to 400 (extremely disturbed). An A index of 30 or
// greater indicates local geomagnetic storm conditions. See Appendix
// B for further explanation.
//
// Missing A- and K-indices are shown as -1.
type Site struct {
	// A is the 24-hour A index
	A int `json:"A"`

	// K is an array of 3-hour K indexes.
	K []int `json:"K"` // For jsonschema
}

// NewSite returns an initialized Site.
func NewSite() *Site {
	return &Site{
		K: make([]int, 8),
	}
}

// Day represents several sites for a given year, month, and day.
type Day struct {
	Year  uint32 `json:"year"`
	Month uint32 `json:"month"`
	Day   uint32 `json:"day"`

	// Sites is map from site name (e.g., "Fredericksburg") to
	// site report.
	Sites map[string]*Site `json:"sites"`

	// Line, when provided, is the string that was parsed to
	// generate this Day's data.
	Line string `json:"line,omitempty"`
}

// ParseDay attempts to parse a line of text into a Day.
func ParseDay(line string) (*Day, error) {
	var r Day
	r.Line = line

	var (
		f = NewSite()
		c = NewSite()
		p = NewSite()
	)

	_, err := fmt.Sscanf(line, "%d %d %d    %02d %02d%02d%02d%02d%02d%02d%02d%02d   %02d %02d%02d%02d%02d%02d%02d%02d%02d   %02d %02d%02d%02d%02d%02d%02d%02d%02d",
		&r.Year, &r.Month, &r.Day,
		&f.A,
		&f.K[0], &f.K[1], &f.K[2], &f.K[3], &f.K[4], &f.K[5], &f.K[6], &f.K[7],
		&c.A,
		&c.K[0], &c.K[1], &c.K[2], &c.K[3], &c.K[4], &c.K[5], &c.K[6], &c.K[7],
		&p.A,
		&p.K[0], &p.K[1], &p.K[2], &p.K[3], &p.K[4], &p.K[5], &p.K[6], &p.K[7])

	if err != nil {
		return nil, err
	}

	r.Sites = map[string]*Site{
		"Fredericksburg": f,
		"College":        c,
		"Planetary":      p,
	}

	return &r, err
}

// ParseFile attempts to read an entire file (including the header)
// and return Day reports.
func ParseFile(s string) ([]*Day, error) {
	head := regexp.MustCompile(HeaderPattern).FindString(s)
	if len(head) == 0 {
		return nil, fmt.Errorf("file has bad header")
	}
	s = s[len(head):]

	acc := make([]*Day, 0, 31)
	for i, line := range strings.SplitN(s, "\n", -1) {
		if len(line) == 0 {
			break
		}
		d, err := ParseDay(line)
		if err != nil {
			return nil, fmt.Errorf("error at line %d: %s\n%s\n", i, err, line)
		}
		acc = append(acc, d)
	}
	return acc, nil
}

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
