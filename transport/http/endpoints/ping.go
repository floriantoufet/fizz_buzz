package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

// Ping does not requires an usecase to optimize code
func (gw *Endpoints) Ping(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "pong")
}
