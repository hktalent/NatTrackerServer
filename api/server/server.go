package server

import (
	"log"
	"net"
	"runtime"
	"sync"
	"time"

	ntp "github.com/hktalent/NatTrackerServer/lib"
)

var (
	bufferPool    sync.Pool
	mq            messageQueue
	nbWorkers     int
	address       string
	UDPPacketSize int = 1500
	maxQueueSize  int = 1000000
)

type messageQueue chan ntp.NatTrackerProtocol

// get message
func (mq messageQueue) enqueue(m ntp.NatTrackerProtocol) {
	mq <- m
}

// do message
func (mq messageQueue) dequeue() {
	for m := range mq {
		handleMessage(m)
		bufferPool.Put([]byte(m.Msg))
	}
}

// 监听
func listenAndReceive(maxWorkers int) error {
	c, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}
	tick1 := time.Tick(time.Duration(time.Second))
	for {
		select {
		case <-tick1:
			{
				if len(mq) < maxQueueSize {
					for i := 0; i < maxWorkers; i++ {
						go mq.dequeue()
						go receive(c)
					}
				}
			}
		}
	}
	return nil
}

// receive accepts incoming datagrams on c and calls handleMessage() for each message
func receive(c net.PacketConn) {
	message := bufferPool.Get().([]byte)
	rlen, remote, err := c.ReadFrom(message[0:])
	if err != nil {
		log.Printf("ReadFrom Error %s", err)
		return
	}
	nntkp := ntp.NewNatTrackerProtocol()
	nntkp.Msg = string(message[0:rlen])
	nntkp.Remote = remote
	nntkp.Conn = c
	nntkp.MsgLen = rlen
	mq.enqueue(*nntkp)
}

// 处理协议
func handleMessage(m ntp.NatTrackerProtocol) {
	ntp.Dispacker(&m)
}

func NewServer(address1 string, nbWorkers1, UDPPacketSize1, maxQueueSize1 int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	address = address1
	nbWorkers = nbWorkers1
	UDPPacketSize = UDPPacketSize1
	maxQueueSize = maxQueueSize1

	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, UDPPacketSize) },
	}
	mq = make(messageQueue, maxQueueSize)
	// log.Println(len(mq))
	listenAndReceive(nbWorkers)
	select {}
}
