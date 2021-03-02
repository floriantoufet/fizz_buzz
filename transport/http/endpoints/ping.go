package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

// Ping will be used for technical purpose
func (gw *Endpoints) Ping(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "pong")
}
