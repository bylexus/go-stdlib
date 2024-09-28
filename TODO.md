# Todo - Alex' golang stdlib

## Router

* [ ] Support for Host matching (with / without regex)
* [ ] Support for sub-routers (nested routers)
* [x] Inject MatchingRoute in request context
* [ ] Support for regex outside the param placeholders

## Middlewares

* [ ] RequestLog: Support for custom format
* [ ] RequestLog: Support for default format: apache log format

## HTTP Server

I want to implement my own HTTP Web Server. Why would anyone do that, you may ask? Here's why:

* I want to understand things. Best way to achieve that is to build it, tinker with it.
* I want to build my own infrastructure for https://alexi.ch as much as possible. This also includes my own web server.
* I am curious of all tech-related things. I want to challenge myself and see if I can do it.
* I want to program more things GO.
* Just for fun!

### What should it support?

I want to start slowly:

* [ ] a small HTTP 1.0 (yes, 1.0) web server: One request, one connection, one thread.
* [ ] standard GO net.http interface support for define handlers / routes
* [ ] support a very limited set of HTTP headers, just enough to deliver some simple web content


Then, in a next step:
* [ ] http 1.1, with multiple requests per connections
* [ ] tls support
* [ ] ... let's see where this journey leads me...

## CLI helper tools

I want to use Go as my go-to language also for small scripts / CLI tools. Therefore I want to establish a set of helper tools for building command line apps.
Some ideas that come to mind:

* [ ] CLI arg parser with flags, commands, subcommands
* text file processing tools:
  * [ ] tools for reading files, based on the new Iterator functions in go 1.23
  * [ ] Read entire file to string
  * [ ] Read file to line array
  * [ ] Read file to lines split by regex
  * [ ] Read file to lines to a specific type / struct (e.g. csv parsing)
  * [ ] Read file line by line with a callback
* Logging tools:
  * make logging to cli and files easy, simple
  * create a logging infrastructure compatible with the go log interface, that supports
	* [ ] file logging
	* [ ] cli logging
	* [ ] combining logs (e.g. "TeeLogger", that logs to file + stdout at the same time)
  * [ ] Output functions for colored output
* Simple facility to run / execute external commands

