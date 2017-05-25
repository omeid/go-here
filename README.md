# Here [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/omeid/go-here)  [![Build Status](https://travis-ci.org/omeid/go-here.svg?branch=master)](https://travis-ci.org/omeid/go-here) [![Go Report Card](https://goreportcard.com/badge/github.com/omeid/go-here)](https://goreportcard.com/report/github.com/omeid/go-here)
Where are you now?


Here is a simple utility package that provides a simple interface to get the location of a Go code source file, from the file itself.

### Why?

I use [go-resources](https://github.com/omeid/go-resources) to embed non-go content with Go binaries. go-resources provides a pattern of using build tags for loading to-be-embeded resources directly from the disk for faster development and iteration.

However, this trick only works if you're embedding resources in your main package and only run it from there, this package is here to help bridge this gap. It will help you embed and live load content of non-main packages as well.


##### Please note that this package exist as a reference and should only be used for in development and debugging purposes, your software binary should _not_ rely on source code location, really.

### Documentation

The package provides only two methods, `Dir() string` which returns the directory of the caller, and `Abs(path string) string` which returns the Absolute path of the provided path relative to the callers directory.

See the [GoDoc](https://godoc.org/github.com/omeid/go-here)



### LICENSE
  [MIT](LICENSE).
