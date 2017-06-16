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
	fmt.Println(network)

	c, err := net.DialTimeout(t.Protocol, network, t.Timeout)
	if err != nil {
		return r, err
	}

	defer c.Close()

	r.RemoteAddr, r.LocalAddr = c.RemoteAddr().String(), c.LocalAddr().String()

	r.Info, err = getTCPInfo(c)
	if err != nil {
		return r, err
	}

	return r, nil
}

func getTCPInfo(c net.Conn) (*tcpinfo.Info, error) {

	tc, err := tcp.NewConn(c)
	if err != nil {
		return nil, err
	}

	var o tcpinfo.Info
	var b [256]byte
	i, err := tc.Option(o.Level(), o.Name(), b[:])
	if err != nil {
		return nil, err
	}

	return i.(*tcpinfo.Info), nil
}
