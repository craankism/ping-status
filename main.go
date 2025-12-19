package main

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	pinger, err := probing.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	ping := pinger.Statistics().AvgRtt // get send/receive/duplicate/rtt stats
	fmt.Println(ping)
}