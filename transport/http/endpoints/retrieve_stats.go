package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

// RetrieveStats returns the FizzBuzz corresponding to most used requests,
// as well as the number of hits for those requests
// returns 200 with JSON response
func (gw *Endpoints) RetrieveStats(w http.ResponseWriter, r *http.Request) {
	// Get stats
	stats, total := gw.uc.RetrieveStats()

	// Build response
	response := map[string]interface{}{
		"total":    total,
		"requests": stats,
	}

	render.JSON(w, r, response)
}
