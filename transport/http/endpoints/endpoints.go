package endpoints

import (
	"fizzbuzz/usecases"
)

type Endpoints struct {
	uc usecases.Usecases
}

func NewEndpoints(uc usecases.Usecases) *Endpoints {
	return &Endpoints{uc: uc}
}
