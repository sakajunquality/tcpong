package ping

import (
	"testing"
	"time"
)

var localTarget = Target{
	Protocol: "tcp",
	Host:     "127.0.0.1",
	Port:     22727,
	Timeout:  time.Duration(1) * time.Second,
}

func TestDialLocal(t *testing.T) {
	ln, err := newLocalListener("tcp")

	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	_, err = localTarget.Dial()

	if err != nil {
		t.Fatal(err)
	}
}
