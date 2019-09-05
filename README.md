# timex - A Go time extension library

[![Build Status](https://www.travis-ci.org/osechet/timex.svg?branch=master)](https://www.travis-ci.org/osechet/timex)
[![codecov](https://codecov.io/gh/osechet/timex/branch/master/graph/badge.svg)](https://codecov.io/gh/osechet/timex)


## Install

```
go get github.com/osechet/timex
```

## Features

The purpose of the library is to provide several extensions to the standard Go time library. For the moment, only GPS time conversion is provided.

## Example


### GPS

```
// Display the GPS time of the current time, in microseconds
fmt.Println(int64(timex.GpsTime(time.Now()).Gps() / time.Microsecond))
```

## Testing

`go test` is used for testing.
