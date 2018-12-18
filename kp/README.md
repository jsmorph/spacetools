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

You'll hopefully see something like:

```JSON
{"year":2018,"month":12,"day":11,"sites":{"College":{"A":12,"K":[0,1,4,5,1,2,2,1]},"Fredericksburg":{"A":7,"K":[1,1,3,3,1,2,2,1]},"Planetary":{"A":7,"K":[1,2,3,2,0,2,2,2]}}}
{"year":2018,"month":12,"day":12,"sites":{"College":{"A":2,"K":[0,0,0,1,2,1,1,0]},"Fredericksburg":{"A":3,"K":[1,1,1,1,1,2,1,0]},"Planetary":{"A":4,"K":[2,1,1,1,1,1,2,1]}}}
{"year":2018,"month":12,"day":13,"sites":{"College":{"A":1,"K":[0,0,0,1,0,1,0,0]},"Fredericksburg":{"A":2,"K":[0,1,0,0,1,2,1,0]},"Planetary":{"A":3,"K":[0,1,0,0,0,1,0,0]}}}
{"year":2018,"month":12,"day":14,"sites":{"College":{"A":0,"K":[0,0,0,1,0,0,0,0]},"Fredericksburg":{"A":3,"K":[1,1,0,1,1,1,1,1]},"Planetary":{"A":3,"K":[1,1,0,1,0,1,1,1]}}}
{"year":2018,"month":12,"day":15,"sites":{"College":{"A":0,"K":[0,0,0,0,0,0,0,0]},"Fredericksburg":{"A":1,"K":[0,0,0,0,1,1,1,0]},"Planetary":{"A":2,"K":[1,0,0,0,0,0,1,0]}}}
{"year":2018,"month":12,"day":16,"sites":{"College":{"A":1,"K":[0,0,1,1,0,0,0,0]},"Fredericksburg":{"A":1,"K":[0,0,1,0,1,1,0,0]},"Planetary":{"A":2,"K":[1,0,1,0,0,1,0,1]}}}
{"year":2018,"month":12,"day":17,"sites":{"College":{"A":-1,"K":[0,0,0,0,2,0,1,-1]},"Fredericksburg":{"A":-1,"K":[1,1,2,1,1,2,2,-1]},"Planetary":{"A":5,"K":[1,1,2,1,1,2,2,-1]}}}
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
