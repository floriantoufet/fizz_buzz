// nolint:unused
// All definition in this package expect for the config object are used to
// ensure config validation. Validation process is made through reflexion so
// unused linter does not see it.
package http

import (
	"errors"
	"fmt"
)

var ErrInvalidPort = errors.New("invalid HTTP port")

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (config Config) GetAddress() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

func (config Config) Check() error {
	if config.Port < 1024 || config.Port > 65535 {
		return ErrInvalidPort
	}

	return nil
}
