# Flect

<p align="center">
<a href="https://godoc.org/github.com/gobuffalo/flect"><img src="https://godoc.org/github.com/gobuffalo/flect?status.svg" alt="GoDoc" /></a>
<a href="https://travis-ci.org/gobuffalo/flect"><img src="https://travis-ci.org/gobuffalo/flect.svg?branch=master" alt="Build Status" /></a>
<a href="https://goreportcard.com/report/github.com/gobuffalo/flect"><img src="https://goreportcard.com/badge/github.com/gobuffalo/flect" alt="Go Report Card" /></a>
</p>

This is a new inflection engine to replace [https://github.com/markbates/inflect](https://github.com/markbates/inflect) designed to be more modular, more readable, and easier to fix issues on than the original.

The `github.com/gobuffalo/flect` package contains "basic" inflection tools, like pluralization, singularization, etc...

The `github.com/gobuffalo/flect/name` package contains more "business" inflection rules like creating proper names, table names, etc...

## Installation

```bash
$ go get -u -v github.com/gobuffalo/flect
```
