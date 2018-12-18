# Simple tool to obtain some ocurrent geomagnetic data

Get [NOAA's recent Space Weather Conditions](https://www.swpc.noaa.gov/products/planetary-k-index).

![current conditions](https://services.swpc.noaa.gov/images/planetary-k-index.gif)

## Building

Install [Go](https://golang.org/doc/install).

Then `go test && go install`

## Usage

```Shell
kp
```

By default, the program gets and
parses
[this file](https://services.swpc.noaa.gov/text/daily-geomagnetic-indices.txt).


If you invoke the program with `kp -h`, you'll see something like

```
Get daily geomagnetic data in JSON

  -a	wrap records in an array
  -d	include lines with records
  -f string
    	filename (if data is local)
  -h	get help
  -s	just print out JSON Schema and stop
  -u string
    	URL for file (default "https://services.swpc.noaa.gov/text/daily-geomagnetic-indices.txt")

```
