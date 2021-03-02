package usecases

type Vanilla struct{}

func NewUsecases() Usecases {
	return &Vanilla{}
}
