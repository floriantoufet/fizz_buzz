package endpoints

import (
	"fizzbuzz/usecases"
)

type Endpoint struct {
	uc usecases.Usecases
}

func NewEndpoint(uc usecases.Usecases) *Endpoint {
	return &Endpoint{uc: uc}
}
