package main

import (
	"flag"
	"log"

	"github.com/hktalent/NatTrackerServer/client"
)

func main() {
	uuid := flag.String("uuid", "", "uuid")
	flag.Parse()
	// get self NAT public ip and port
	x1 := client.Nat()
	log.Println("your Nat Public Ip and Port: ", x1)
	// create you uuid, on other client use the uuid, discover each other
	if *uuid == "" {
		*uuid = client.GenerateUUID()
	}
	log.Println("your uuid: ", *uuid)

	// reg self NAT public ip and port,and get other member lists
	a := client.AutoRegSelf(*uuid, "")
	log.Println("first is your ip, member is:")
	for _, x := range a {
		log.Println(x)
	}

}
