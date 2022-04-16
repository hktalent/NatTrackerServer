package lib

import (
	"fmt"
	"regexp"
	"strings"
)

var SvKvDbOp *KvDbOp = NewKvDbOp()

// uuid len >= 10
// mac addres format:
// 10:1c:2a:f2:49:92
// 101c2af24992
func Dispacker(cc1 *NatTrackerProtocol) {
	// //51pwn/P2P&E2E/[uuid]/your_publicIpPort_or_0/your_LanIps/self_mac_Addres
	aReg := []*regexp.Regexp{regexp.MustCompile(`^\/\/51pwn\/P2P&E2E\/[^\/]{10,}\/[^\/]+/[^\/]+/[^\/]+`)}

	defer cc1.Conn.Close()
	switch {
	case Nat == cc1.Msg:
		{
			s := fmt.Sprintf(cc1.Remote.String())
			// log.Println(s)
			cc1.Conn.WriteToUDP([]byte(s), cc1.Remote)
		}
	case strings.HasPrefix(cc1.Msg, NatPrefix):
		{
			aData := strings.Split(cc1.Msg[len(NatPrefix):], "/")
			// uuid,public ip port,sefl mac ip
			if 3 == len(aData) {

			}
		}
	case strings.HasPrefix(cc1.Msg, SzTrackerS2SPrefix):
		{
			aData := strings.Split(cc1.Msg[len(SzTrackerS2SPrefix):], "/")
			switch {
			case aReg[0].MatchString(cc1.Msg):
				{

				}
			}
		}

	}

}
