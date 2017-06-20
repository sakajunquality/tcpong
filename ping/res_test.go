package ping

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringTCP(t *testing.T) {
	res := Res{
		protocol:   "tcp",
		seq:        0,
		remoteAddr: "127.0.0.1:53",
		localAddr:  "127.0.0.1:0",
		state:      "established",
		rtt:        time.Duration(3) * time.Millisecond,
	}

	expected := "tcp_seq=0 state=established from=127.0.0.1:53 rtt=3ms"
	actual := res.String()

	assert.Equal(t, expected, actual)
}

func TestStringNonTCP(t *testing.T) {
	res := Res{
		protocol:   "udp",
		seq:        0,
		remoteAddr: "127.0.0.1:500",
		localAddr:  "127.0.0.1:0",
	}

	expected := "udp_seq=0 state= from=127.0.0.1:500 rtt=0s"
	actual := res.String()

	assert.Equal(t, expected, actual)
}
