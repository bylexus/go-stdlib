# bylexus' golang stdlib

This is my personal golang stdlib, containing various cross-project library functions
and data structures.

## Note on naming convention

All my packages start with an "e" (e.g. for "enhanced", or "extended". Another interpretation may be "evil"?),
to avoid confusion with the standard library, and to avoid confusion with the developer's mindset.

## Available Functions

* `eerr.PanicOnErr` takes an error value, and, if not nil, outputs it and panics.
* `elog.SeverityLogger` implements a struct with loggers for different severities
* `emaps.GetMapKeys` returns a slice of all keys of the given map
* `efunctional.Filter` filters a slice by a given predicate function
* `efunctional.Map` Applies a function to each element of a slice, returning a new slice with the results.
* `efunctional.Reduce` applies a function to each element of a slice and an accumulator value, returning a single reduced value
* `eslices.InSlice` checks if a given value is in a slice
* `estrings.SplitRe` splits a string by a regex

## Available tools

* `ehttp.Router` offers a `http.Handler` which supports better routing than the default `http.ServeMux` routing mechanism
* http middlewares
	* `ehttp.middleware.ClientLimit`: Limits the number of concurrent requests
	* `ehttp.middleware.Delay`: Delays the processing of a single request for a given amount of time
	* `ehttp.middleware.RequestLog`: Logs the request to the given logger in a pre-defined format
	* `ehttp.middleware.HtmlContent`: Adds the 'Content-Type: text/html' content type header to the output

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
