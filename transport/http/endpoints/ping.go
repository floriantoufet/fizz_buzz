package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (gw *Endpoint) Ping(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, gw.uc.Ping())
}
