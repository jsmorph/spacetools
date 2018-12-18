package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseDay(t *testing.T) {
	line := `2018 10 01     6  0 1 1 2 3 2 1 2    19  0 1 2 4 6 4 2 1     9  0 2 1 2 4 2 2 2`
	s, err := ParseDay(line)
	if err != nil {
		t.Fatal(err)
	}

	js, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", js)
}

func TestFile(t *testing.T) {
	s := `:Product: Daily Geomagnetic Data     quar_DGD.txt
:Issued: 2130 UT 17 Dec 2018
#
#  Prepared by the U.S. Dept. of Commerce, NOAA, Space Weather Prediction Center
#  Please send comment and suggestions to SWPC.Webmaster@noaa.gov
#
#             Current Quarter Daily Geomagnetic Data
#
#
#                Middle Latitude        High Latitude            Estimated
#              - Fredericksburg -     ---- College ----      --- Planetary ---
#  Date        A     K-indices        A     K-indices        A     K-indices
2018 10 01     6  0 1 1 2 3 2 1 2    19  0 1 2 4 6 4 2 1     9  0 2 1 2 4 2 2 2
2018 10 02     6  3 2 2 1 2 1 1 1     4  2 1 2 1 1 1 1 0     8  4 3 2 1 2 1 2 0
2018 10 03     4  0 1 0 1 1 2 2 2     2  0 0 1 0 0 1 1 2     6  1 1 1 2 1 1 2 3
2018 10 04     4  1 0 1 0 2 2 2 1     4  1 0 0 0 2 3 1 1     4  1 0 1 0 1 2 1 1
2018 10 05     8  3 2 2 1 3 1 1 2     6  1 2 1 3 2 1 1 1     9  3 3 2 2 3 1 1 2
2018 10 06     5  3 1 2 2 1 0 1 0     6  1 1 2 4 2 0 0 0     6  3 1 2 2 1 0 1 0
2018 10 07    13  0 0 2 3 3 4 4 3    44  0 0 2 6 7 6 4 3    24  0 1 1 3 4 5 5 5
2018 10 08    16  4 4 3 3 3 3 1 2    53  3 5 6 6 7 5 2 1    21  4 4 4 3 4 3 2 2
2018 10 09    14  2 3 3 3 3 3 3 3    39  2 3 5 6 6 5 3 3    17  3 3 3 4 3 4 3 3
2018 10 10    12  3 1 1 2 3 3 4 2    33  2 1 1 6 6 5 4 3    18  4 1 1 2 3 3 5 3
2018 10 11     7  3 2 2 1 2 1 2 1    13  2 2 4 3 4 2 2 1     9  3 3 3 2 2 2 2 1
2018 10 12     4  1 2 2 1 1 1 1 1    10  1 1 3 5 2 1 0 1     5  1 2 2 1 1 1 1 1
2018 10 13    12  1 2 1 1 2 4 3 4    15  1 1 1 0 1 6 3 3    14  1 3 1 1 2 5 4 4
2018 10 14     4  1 1 1 1 2 2 1 1     8  2 1 1 4 3 1 0 1     6  2 1 1 2 2 2 1 2
`
	ds, err := ParseFile(s)
	if err != nil {
		t.Fatal(err)
	}
	for i, d := range ds {
		js, err := json.Marshal(d)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%02d %s\n", i, js)
	}
}
