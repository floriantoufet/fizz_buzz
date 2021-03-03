package endpoints

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"fizzbuzz/usecases"
)

// FizzBuzz returns a string with numbers from 1 to limit where:
// all multiples of request.FizzModulo are replaced by request.FizzString,
// all multiples of request.BuzzModulo are replaced by request.BuzzString,
// all multiples of request.FizzModulo and request.BuzzModulo are replaced by FizzStringBuzzString
// returns 200 with fizzBuzz string if success
// returns 404 if modulus and limit are not present, not int, zero or negatives
// returns 500 if get an unexpected error
func (gw *Endpoints) FizzBuzz(w http.ResponseWriter, r *http.Request) {
	// Get request params
	var (
		fizzModulo, buzzModulo, limit *int
		fizzString, buzzString        *string
	)

	if value, err := strconv.Atoi(r.URL.Query().Get("fizz_modulo")); err == nil {
		fizzModulo = &value
	}
	if value, err := strconv.Atoi(r.URL.Query().Get("buzz_modulo")); err == nil {
		buzzModulo = &value
	}
	if value, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = &value
	}

	if value := r.URL.Query().Get("fizz_string"); value != "" {
		fizzString = &value
	}
	if value := r.URL.Query().Get("buzz_string"); value != "" {
		buzzString = &value
	}

	// Get fizzBuzz string
	fizzBuzzResponse, err := gw.uc.FizzBuzz(fizzModulo, buzzModulo, limit, fizzString, buzzString)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrInvalidLimit), errors.Is(err, usecases.ErrInvalidModulo):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.PlainText(w, r, fizzBuzzResponse)
}
