package utils

import (
	"bytes"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Convert Lambda request → http.Request
func ConvertToHTTPRequest(req events.LambdaFunctionURLRequest) (*http.Request, error) {

	body := io.NopCloser(bytes.NewBufferString(req.Body))

	httpReq, err := http.NewRequest(
		req.RequestContext.HTTP.Method,
		req.RawPath,
		body,
	)
	if err != nil {
		return nil, err
	}

	// Add query params
	q := httpReq.URL.Query()
	for k, v := range req.QueryStringParameters {
		q.Add(k, v)
	}
	httpReq.URL.RawQuery = q.Encode()

	// Add headers
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	return httpReq, nil
}

type ResponseRecorder struct {
	Body       bytes.Buffer
	StatusCode int
	HeaderMap  http.Header
}

func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		StatusCode: 200,
		HeaderMap:  make(http.Header),
	}
}

func (r *ResponseRecorder) Header() http.Header {
	return r.HeaderMap
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	return r.Body.Write(b)
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
}
