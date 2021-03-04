package endpoints

import (
	"github.com/floriantoufet/fizzbuzz/usecases"
)

type Endpoints struct {
	uc usecases.Usecases
}

func NewEndpoints(uc usecases.Usecases) *Endpoints {
	return &Endpoints{uc: uc}
}
