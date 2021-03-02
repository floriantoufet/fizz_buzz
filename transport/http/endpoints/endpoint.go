package endpoints

import (
	"fiz_buz/usecases"
)

type Endpoint struct {
	uc usecases.Usecases
}

func NewEndpoint(uc usecases.Usecases) *Endpoint {
	return &Endpoint{uc: uc}
}
