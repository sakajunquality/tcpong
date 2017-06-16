package ping

import (
	"net"
	"strconv"
	"testing"
	"time"
)

func TestDialLocal(t *testing.T) {
	ln, err := newLocalListener("tcp")

	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	var localTarget = Target{
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Timeout:  time.Duration(1) * time.Second,
	}

	_, port, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		t.Fatal(err)
	}

	localTarget.Port = intPort
	fmt.println(intPort)

	_, err = localTarget.Dial()

	if err != nil {
		t.Fatal(err)
	}
}
