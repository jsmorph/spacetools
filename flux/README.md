# Simple tool to extact some GEOS data as CSV

See
[GOES Space Environment Monitor Data](https://www.ngdc.noaa.gov/stp/satellite/goes/) for
background, and
see
[this directory](https://satdat.ngdc.noaa.gov/sem/goes/data/new_avg/2018/12/goes15/csv) for
some actual files that this program can parse.

The output is CSV that's simpler than the input CSV.

## Building

Install [Go](https://golang.org/doc/install).

Then `go install`

## Usage

```Shell
wget https://satdat.ngdc.noaa.gov/sem/goes/data/new_avg/2018/12/goes15/csv/g15_epead_cpflux_5m_20181201_20181231.csv
flux -f g15_epead_cpflux_5m_20181201_20181231.csv -w time_tag,ZPGT1E | head
```

Example:

```
flux -f g15_epead_cpflux_5m_20181201_20181231.csv -w time_tag,ZPGT1E,ZPGT5E | head
time_tag,ZPGT1E,ZPGT5E
2018-12-01 00:00:00.000,1.4299E+01,3.7709E-01
2018-12-01 00:05:00.000,1.4455E+01,1.7297E-01
2018-12-01 00:10:00.000,1.5017E+01,3.7720E-01
2018-12-01 00:15:00.000,1.8807E+01,2.4295E-01
2018-12-01 00:20:00.000,1.3551E+01,1.7232E-01
2018-12-01 00:25:00.000,1.3996E+01,1.7232E-01
2018-12-01 00:30:00.000,2.0293E+01,3.0038E-01
2018-12-01 00:35:00.000,1.4956E+01,1.8552E-01
2018-12-01 00:40:00.000,1.9875E+01,3.0038E-01
```

