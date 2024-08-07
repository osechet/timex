# timex - A Go time extension library

[![Build Status](https://github.com/osechet/timex/actions/workflows/go.yml/badge.svg)](https://github.com/osechet/timex/actions)
[![codecov](https://codecov.io/gh/osechet/timex/branch/master/graph/badge.svg)](https://codecov.io/gh/osechet/timex)

## Install

```sh
go get github.com/osechet/timex
```

## Features

The purpose of the library is to provide several extensions to the standard Go time library. For the moment, only GPS time conversion is provided.

## Example

### GPS

```go
// Display the GPS time of the current time, in microseconds
fmt.Println(int64(timex.GpsTime(time.Now()).Gps() / time.Microsecond))
```

## Testing

`go test` is used for testing.
