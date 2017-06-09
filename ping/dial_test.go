package ping

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var localTarget = Target{
	Protocol: "tcp",
	Host:     "127.0.0.1",
	Port:     0,
	Timeout:  time.Duration(1) * time.Second,
}

func TestDialLocal(t *testing.T) {
	ln, err := newLocalListener("tcp")

	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	r, err := localTarget.Dial()

	if err != nil {
		t.Fatal(err)
	}

	expected := localTarget.Host
	actual := r.RemoteAddr

	assert.Equal(t, expected, actual)
}
