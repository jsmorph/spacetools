// Package main extracts GEOS space weather data.
//
// See https://www.ngdc.noaa.gov/stp/satellite/goes/ for background
// and
// https://satdat.ngdc.noaa.gov/sem/goes/data/new_avg/2018/12/goes15/csv
// for some actual files that this program can parse.
//
// The output is CSV that's simpler than the input CSV.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		filename = flag.String("f", "g15_epead_cpflux_5m_20181201_20181231.csv", "filename to read")
		wanted   = flag.String("w", "time_tag,ZPGT1E,ZPGT5E", "fields to extract")
		help     = flag.Bool("h", false, "get help")
	)

	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Extract some GEOS data as CSV\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Example:

  wget https://satdat.ngdc.noaa.gov/sem/goes/data/new_avg/2018/12/goes15/csv/g15_epead_cpflux_5m_20181201_20181231.csv
  flux -f g15_epead_cpflux_5m_20181201_20181231.csv -w time_tag,ZPGT1E

`)
		return
	}

	bs, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	want := strings.SplitN(*wanted, ",", -1)

	scanner := bufio.NewScanner(bytes.NewReader(bs))

	data := false
	var cols map[string]int

	for i, col := range want {
		if 0 < i {
			fmt.Printf(",")
		}
		fmt.Printf("%s", col)
	}
	fmt.Printf("\n")

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if !data {
			if line == "data:" {
				data = true
			}
			continue
		}
		if cols == nil {
			cols = make(map[string]int, 128)
			for i, col := range strings.SplitN(line, ",", -1) {
				cols[col] = i
			}
			continue
		}
		xs := strings.SplitN(line, ",", -1)
		// ToDo: Check lengths.

		for i, col := range want {
			at := cols[col]
			x := xs[at]
			if 0 < i {
				fmt.Printf(",")
			}

			fmt.Printf("%s", x)
		}
		fmt.Printf("\n")

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
