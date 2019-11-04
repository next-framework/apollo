package router

import (
	"github.com/next-frmework/apollo"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
	a *apollo.Apollo
}
