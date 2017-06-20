package ping

import (
	"fmt"
	"net"

	"github.com/mikioh/tcp"
	"github.com/mikioh/tcpinfo"
)

func (t *Target) Dial() (Res, error) {
	r := Res{protocol: t.Protocol, seq: t.Seq}
	t.Seq++
	network := fmt.Sprintf("%s:%d", t.Host, t.Port)

	c, err := net.DialTimeout(t.Protocol, network, t.Timeout)
	if err != nil {
		return r, err
	}

	defer c.Close()
	r.remoteAddr, r.localAddr = c.RemoteAddr().String(), c.LocalAddr().String()

	switch t.Protocol {
	case "tcp":
		info, err := getTCPInfo(c)
		if err != nil {
			return r, err
		}
		r.rtt = info.RTT
		r.state = fmt.Sprintf("%s", info.State)
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
