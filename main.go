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
	address := pinger.Statistics().Addr
	ping := pinger.Statistics().AvgRtt // get send/receive/duplicate/rtt stats
	packetLoss := pinger.Statistics().PacketLoss
	TTL := pinger.Statistics().TTLs
	fmt.Printf("%-15s %s\n", "Website:", address)
	fmt.Printf("%-15s %d ms\n", "Ping:", ping.Milliseconds())
	fmt.Printf("%-15s %.2f%%\n", "Packet Loss:", packetLoss)
	fmt.Printf("%-15s %v\n", "TTLs received:", TTL)
}
