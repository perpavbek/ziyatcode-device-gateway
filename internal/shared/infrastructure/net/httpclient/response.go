package httpclient

import "net/http"

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}