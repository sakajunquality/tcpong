package ping

import (
	"fmt"
	"net"
)

func newLocalListener(network string) (net.Listener, error) {
	switch network {
	case "tcp":
		if ln, err := net.Listen("tcp4", ":22727"); err == nil {
			return ln, nil
		}
	}
	return nil, fmt.Errorf("%s is not supported", network)
}
