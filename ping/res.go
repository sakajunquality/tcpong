package ping

import (
	"fmt"
	"time"
)

type Res struct {
	protocol   string
	seq        int
	remoteAddr string
	localAddr  string
	state      string
	rtt        time.Duration
}

func (r *Res) String() string {
	return fmt.Sprintf("%s_seq=%d state=%s from=%s rtt=%s", r.protocol, r.seq, r.state, r.remoteAddr, r.rtt)
}
