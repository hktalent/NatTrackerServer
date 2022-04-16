package main

import (
	"flag"
	"runtime"

	"github.com/hktalent/NatTrackerServer/api/server"
)

var (
	nbWorkers     int
	address       string
	UDPPacketSize int = 1500
	maxQueueSize  int = 1000000
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&address, "addr", ":43283", "Address of the UDP server to test")
	flag.IntVar(&nbWorkers, "concurrency", runtime.NumCPU(), "Number of workers to run in parallel")
	flag.IntVar(&UDPPacketSize, "UDPPacketSize", 1500, "Number of UDPPacketSize")
	flag.IntVar(&maxQueueSize, "maxQueueSize", 1000000, "Number of maxQueueSize")

	flag.Parse()

	server.NewServer(address, nbWorkers, UDPPacketSize, maxQueueSize)
}
