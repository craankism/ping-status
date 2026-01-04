package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func pinger(domains string) (ping float64, packetLoss float64) {
	//icmp echo
	pinger, err := probing.NewPinger(domains)
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	ping = float64(pinger.Statistics().AvgRtt)
	packetLoss = pinger.Statistics().PacketLoss
	return ping, packetLoss
}

func main() {
	//env var for domains
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	domains := os.Getenv("DOMAINS")

	//prometheus
	//set new gauges
	ping := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:      "ping",
		Subsystem: domains,
		Namespace: "craankism",
		Help:      "Checking ping of server in ms",
	},
		func() float64 {
			ping, _ := pinger(domains)
			return ping
		},
	)
	packetLoss := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:      "packet_loss",
		Subsystem: domains,
		Namespace: "craankism",
		Help:      "Checking packet loss of server in %",
	},
		func() float64 {
			_, packetLoss := pinger(domains)
			return packetLoss
		},
	)
	//register gauges
	prometheus.MustRegister(ping)
	prometheus.MustRegister(packetLoss)

	//remove go collector
	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
