package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

// RetrieveStats returns the FizzBuzz corresponding to most used requests,
// as well as the number of hits for those requests
// returns 200 with JSON response
func (gw *Endpoints) ResetStats(w http.ResponseWriter, r *http.Request) {
	// Delete stats
	gw.uc.ResetStats()

	render.JSON(w, r, "success")
}
