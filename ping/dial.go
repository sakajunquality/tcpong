package ping

import (
	"fmt"
	"net"

	"github.com/mikioh/tcp"
	"github.com/mikioh/tcpinfo"
)

type Res struct {
	RemoteAddr string
	LocalAddr  string
	Info       *tcpinfo.Info
}

func (t *Target) Dial() (Res, error) {
	r := Res{}
	network := fmt.Sprintf("%s:%d", t.Host, t.Port)
	c, err := net.DialTimeout(t.Protocol, network, t.Timeout)

	if err != nil {
		return r, err
	}

	defer c.Close()

	r.RemoteAddr = c.RemoteAddr().String()
	r.LocalAddr = c.LocalAddr().String()

	tc, err := tcp.NewConn(c)
	if err != nil {
		return r, err
	}

	var o tcpinfo.Info
	var b [256]byte
	i, err := tc.Option(o.Level(), o.Name(), b[:])
	if err != nil {
		return r, err
	}

	r.Info = i.(*tcpinfo.Info)

	return r, nil
}
