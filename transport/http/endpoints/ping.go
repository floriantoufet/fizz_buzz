package endpoints

import (
	"net/http"
)

func (gw *Endpoint) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(gw.uc.Ping()))
}
