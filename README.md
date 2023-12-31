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
