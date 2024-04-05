package httprequestpath

import (
	"net/http"
	"net/url"
)

// HTTPRequestPath is given a *http.Request and it returns the path from the HTTP-request.
//
// It deals with the error handling that should usually never happen — but weird things
// (such as things being nil that usually are never nil) can happen when a server is
// severely stressed, and thus you should check for if you want your software to handle
// that type of server stress more gracefully.
func HTTPRequestPath(request *http.Request) (string, error) {
	if nil == request {
		return "", errNilHTTPRequest
	}

	var urlocator *url.URL = request.URL
	if nil == urlocator {
		return "", errNilHTTPRequestURI
	}

	var path string = urlocator.Path

	return path, nil
}
