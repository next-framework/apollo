package apollo

import "net/http"

type Model interface {
	ResponseWriter() http.ResponseWriter

	ResetResponseWriter(w http.ResponseWriter)

	Request() *http.Request

	ResetRequest(r *http.Request)

	PathVariables() PathVariables
}
