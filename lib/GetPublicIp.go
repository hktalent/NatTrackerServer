package lib

import (
	"net"
)

func GetPublicIP(szServerIp string) (string, error) {

	address := szServerIp
	conn, err := net.Dial("udp4", address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte("nat"))
	if nil != err {
		return "", err
	} else {
		p := make([]byte, 2048)

		n, err := conn.Read(p)
		if err == nil {
			return string(p[0:n]), nil
		} else {
			return "", err
		}
	}
}
