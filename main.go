package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sakajunquality/tcpong/ping"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tcpong"
	app.Usage = "Ping over TCP"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "i",
			Value: 1,
			Usage: "Request Interval (Sec)",
		},
		cli.IntFlag{
			Name:  "t",
			Value: 1,
			Usage: "Request Timeout (Sec)",
		},
		cli.StringFlag{
			Name:  "p",
			Value: "tcp",
			Usage: "Protocol",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() != 2 {
			return cli.NewExitError("illegal number of args", 2)
		}

		ch := make(chan string, 1)
		var seq int

		port, err := strconv.Atoi(c.Args().Get(1))
		if err != nil {
			return cli.NewExitError("Port number should be interger", 2)
		}

		reqInterval, reqTimeout := time.Duration(c.GlobalInt("i"))*time.Second, time.Duration(c.GlobalInt("t"))*time.Second
		reqProtocol := strings.ToLower(c.GlobalString("p"))

		t := ping.Target{
			Protocol: reqProtocol,
			Host:     c.Args().Get(0),
			Port:     port,
			Timeout:  reqTimeout,
		}

		if !t.IsValid() {
			return cli.NewExitError("Input values are not valid", 2)
		}

		color.Blue("%s PING %s:%d Interval:%s Timeout:%s", t.Protocol, t.Host, t.Port, reqInterval, reqTimeout)

		for {
			go func() {
				r, err := t.Dial()
				seq++

				if err != nil {
					ch <- fmt.Sprintf("Error: %s", err)
					return
				}

				ch <- fmt.Sprintf("state=%s fromr=%s rtt=%s", r.Info.State, r.RemoteAddr, r.Info.RTT)
			}()

			fmt.Printf("tcp_seq=%d %s\n", seq, <-ch)
			time.Sleep(reqInterval)
		}
	}

	app.Run(os.Args)
}
