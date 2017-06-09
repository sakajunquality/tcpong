package ping

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

var validTarget = Target{
	Protocol: "tcp",
	Host:     "8.8.8.8",
	Port:     53,
	Timeout:  time.Duration(1) * time.Second,
}

func TestValidTarget(t *testing.T) {
	actual := validTarget.IsValid()
	assert.True(t, actual)
}

func TestInvalidPort(t *testing.T) {
	validTarget.Port = 1234567

	actual := validTarget.IsValid()
	assert.False(t, actual)
}
