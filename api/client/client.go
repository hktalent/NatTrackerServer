package client

import (
	"log"

	"github.com/hktalent/NatTrackerServer/lib"
)

var NatClient = lib.NewNatTrackerProtocol()

// suport IPv4 and IPv6
func Nat() string {
	return NatClient.GetNatPublicIpAndPort4Client()
}

func GetAllTrackerServer() []string {
	return NatClient.GetAllTrackerServer()
}

// Generate UUID
func GenerateUUID() string {
	return NatClient.GenerateUUID()
}

// auto get all tracker server, and pubsh memory
// reg self
// get all member lists
// if uuid = "", auto generate uuid,and print send data
// why have selfPubIpPort?If you have already monitored the port and have a public Internet IP,
//    maybe you need to directly expose the IP and port to the tracker,
//    so that the first one returned by the tracker is no longer Nat's address, but the address given by yourself.
func AutoRegSelf(uuid, selfPubIpPort string) []string {
	bPrt := false
	if "" == uuid {
		uuid = GenerateUUID()
		bPrt = true
	}
	if "" == selfPubIpPort {
		selfPubIpPort = NatClient.GetPublicIP2Line("")
	}
	s := lib.NatPrefix + uuid + "/" + selfPubIpPort + "/" + NatClient.GetMacIPAddrLists2Line4Client()
	if bPrt {
		log.Println("your request data is: ", s)
	}
	return NatClient.RegAndGetAllMemberLists4Client(s)
}
