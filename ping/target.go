package ping

import (
	"time"

	"gopkg.in/validator.v2"
)

type Target struct {
	Protocol string        `validate:"nonzero, regexp=^(tcp|udp)$"`
	Host     string        `validate:"nonzero"`
	Port     int           `validate:"min=0, max=65535"`
	Timeout  time.Duration `validate:"min=0"`
}

func (t *Target) IsValid() bool {
	if errs := validator.Validate(t); errs != nil {
		return false
	}
	return true
}
