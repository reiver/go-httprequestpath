# go-httprequestpath

Package httprequestpath provides tools for getting the __path__ from an `http.Request`, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-httprequestpath

[![GoDoc](https://godoc.org/github.com/reiver/go-httprequestpath?status.svg)](https://godoc.org/github.com/reiver/go-httprequestpath)

## Example

```
import "github.com/reiver/go-httprequestpath"

// ...


path, err := httprequestpath.HTTPRequestPath(request)

if nil != err {
	var code int = http.StatusInternalServerError
	var text string = http.StatusText(code)

	http.Error(responseWriter, text, code)
	return
}
```

## Import

To import package **httprequestpath** use `import` code like the follownig:
```
import "github.com/reiver/go-httprequestpath"
```

## Installation

To install package **httprequestpath** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-httprequestpath
```

## Author

Package **httprequestpath** was written by [Charles Iliya Krempeaux](http://reiver.link)
