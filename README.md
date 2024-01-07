# bylexus' golang stdlib

This is my personal golang stdlib, containing various cross-project library functions
and data structures.

## Available Functions

* `err.PanicOnErr` takes an error value, and, if not nil, outputs it and panics.
* `log.SeverityLogger` implements a struct with loggers for different severities
* `maps.GetMapKeys` returns a slice of all keys of the given map
* `slices.Filter` filters a slice by a given predicate function
* `slices.InSlice` checks if a given value is in a slice
* `strings.SplitRe` splits a string by a regex

## Available tools

* `Router` offers a `http.Handler` which supports better routing than the default `http.ServeMux` routing mechanism
* http middlewares
	* `http.middleware.ClientLimit`: Limits the number of concurrent requests
	* `http.middleware.Delay`: Delays the processing of a single request for a given amount of time
	* `http.middleware.RequestLog`: Logs the request to the given logger in a pre-defined format
	* `http.middleware.HtmlContent`: Adds the 'Content-Type: text/html' content type header to the output

## Run tests

Use the prepared shell script `run-tests.sh` to run the tests.

If you want to run specific tests, call them either by single package:

```sh
go test -v github.com/bylexus/go-stdlib/threads
```

or by using the `...` sub-directory matcher:

```sh
# all:
go test -v ./...
# single subdir:
go test -v ./http/...
```
