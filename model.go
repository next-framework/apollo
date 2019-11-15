package apollo

import "net/http"

type PathVariables struct {
	Storage
}

type Model interface {
	ResponseWriter() http.ResponseWriter

	ResetResponseWriter(w http.ResponseWriter)

	Request() *http.Request

	ResetRequest(r *http.Request)

	PathVariables() PathVariables

	Attributes() Storage

	Context() Context
}
