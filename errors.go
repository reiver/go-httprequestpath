package httprequestpath

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilHTTPRequest    = erorr.Error("nil HTTP-request")
	errNilHTTPRequestURI = erorr.Error("nil HTTP-request-URI")
)

