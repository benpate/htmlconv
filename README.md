# HTMLConv 🤓

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/htmlconv)
[![Build Status](https://img.shields.io/github/workflow/status/benpate/htmlconv/Go/main)](https://github.com/benpate/htmlconv/actions/workflows/go.yml)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/htmlconv.svg?style=flat-square)](https://codecov.io/gh/benpate/htmlconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/htmlconv?style=flat-square)](https://goreportcard.com/report/github.com/benpate/htmlconv)
[![Version](https://img.shields.io/github/v/release/benpate/htmlconv?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/htmlconv/releases)

## Simple Go (golang) library for manipulating HTML content

```go
FromText(text) html

ToText(html) text

IsHTML(string) bool

RemoveTags(html) string

Summary(html) string

CollapseWhitespace(string) string
```

## Pull Requests Welcome

Original versions of this library have been used in production on commercial applications for years, and the extra data collection has been a tremendous help for everyone involved.

I'm now open sourcing this library, and others, with hopes that you'll also benefit from a more robust error package.

Please use GitHub to make suggestions, pull requests, and enhancements.  We're all in this together! 🤪
