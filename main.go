package main

import (
	"fmt"
	"os"
	"time"

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

		hostString, portString := c.Args().Get(0), c.Args().Get(1)
		// to do validate

		ch := make(chan string, 1)
		var seq int

		t := ping.Target{
			Protocol: c.GlobalString("p"),
			Host:     hostString,
			Port:     portString,
			Timeout:  time.Duration(c.GlobalInt("t")) * time.Second,
		}

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
			time.Sleep(time.Duration(c.GlobalInt("i")) * time.Second)
		}
	}

	app.Run(os.Args)
}
