package endpoints

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"fizzbuzz/domains"
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
	fizzModulo, err := strconv.Atoi(r.URL.Query().Get("fizz_modulo"))
	if err != nil {
		http.Error(w, "expect int for fizz_modulo", http.StatusBadRequest)
		return
	}
	buzzModulo, err := strconv.Atoi(r.URL.Query().Get("buzz_modulo"))
	if err != nil {
		http.Error(w, "expect int for buzz_modulo", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "expect int for limit", http.StatusBadRequest)
		return
	}

	fizzBuzzRequest := domains.FizzBuzz{
		FizzModulo: fizzModulo,
		BuzzModulo: buzzModulo,
		Limit:      limit,
		FizzString: r.URL.Query().Get("fizz_string"),
		BuzzString: r.URL.Query().Get("buzz_string"),
	}

	// Get fizzBuzz string
	fizzBuzzResponse, err := gw.uc.FizzBuzz(fizzBuzzRequest)
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
