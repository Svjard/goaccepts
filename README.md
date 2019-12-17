# GoAccepts

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/fc86e14746f04abe966d3c771cfb41af)](https://app.codacy.com/app/Svjard/goaccepts?utm_source=github.com&utm_medium=referral&utm_content=Svjard/goaccepts&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/Svjard/goaccepts)](https://goreportcard.com/report/github.com/Svjard/goaccepts) [![Build Status](https://travis-ci.com/Svjard/goaccepts.svg?branch=master)](https://travis-ci.com/Svjard/goaccepts) [![codecov](https://codecov.io/gh/Svjard/goaccepts/branch/master/graph/badge.svg)](https://codecov.io/gh/Svjard/goaccepts)

Utility library that can parse the `Accept*` header to pull out encodings/languages/charsets/mimetypes for use by an API.

## Usage

```
import "github.com/Svjard/goaccepts"

// Header sent in request with "en-US, it;q=0.6"
...
goaccepts.Languages(res.headers["Accept-Language"])

// will return []string{"en-US","it"}
```

## Development

Each header type is seperated into its own file and parser.

Accept          - media-type.go  
Accept-Encoding - encoding.go  
Accept-Charset  - charset.go  
Accept-Language - language.go   
  
RFC Reference: https://tools.ietf.org/html/rfc2616 . 
  
Each module returns has two specific functions, the first returns a raw string array representation of the accepted entities in order of weight. The second returns an array of structures which provides a specific breakdown and detail about each parsed entity.  
  
## Testing

To run the test suite run the following command:

```go test -v ./...```

To run the code coverage report run the following command:

```./go.test.sh```
