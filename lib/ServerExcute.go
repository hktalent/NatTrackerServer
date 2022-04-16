package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var SvKvDbOp *KvDbOp = NewKvDbOp()

// uuid len >= 10
// mac addres format:
// 10:1c:2a:f2:49:92
// 101c2af24992
func Dispacker(cc1 *NatTrackerProtocol) {
	// //51pwn/P2P&E2E/[uuid]/your_publicIpPort_or_0/your_LanIps/self_mac_Addres
	// aReg := []*regexp.Regexp{regexp.MustCompile(`^\/\/51pwn\/P2P&E2E\/[^\/]{10,}\/[^\/]+/[^\/]+/[^\/]+`)}

	// defer cc1.Conn.Close()
	// log.Println(cc1.Msg, NatPrefix)
	switch {
	case Nat == cc1.Msg:
		{
			s := fmt.Sprintf(cc1.Remote.String())
			cc1.Conn.WriteTo([]byte(s), cc1.Remote)
		}
	case strings.HasPrefix(cc1.Msg, NatPrefix):
		{
			aData := strings.Split(cc1.Msg[len(NatPrefix):], "/")
			// log.Println(cc1.Msg, len(aData), NatPrefix)
			// uuid,public ip port,sefl mac ip
			if 3 == len(aData) {
				if aData[1] == "0" {
					aData[1] = cc1.Remote.String()
				}
				abR, err := SvKvDbOp.Get(aData[0])
				aRstData := ""
				// []string{aData[1] + SzLanIpSep + aData[2]}
				// 获取到数据的情况
				if nil == err && nil != abR && 0 < len(abR) {
					aRstData = string(abR)
				}
				szCur := aData[1] + SzLanIpSep + aData[2]
				if "" != aRstData {
					if -1 < strings.Index(aRstData, szCur) {
						aRstData = strings.Replace(aRstData, szCur, "", -1)
					}
					aRstData = szCur + "\n" + aRstData
				} else {
					aRstData = szCur
				}
				aRstData = strings.Replace(aRstData, "\n\n", "\n", -1)
				aBts := []byte(aRstData)
				cc1.Conn.WriteTo(aBts, cc1.Remote)
				SvKvDbOp.Put(aData[0], aBts)
			}
		}
	case strings.HasPrefix(cc1.Msg, SzTrackerS2SPrefix):
		{
			// for client get all Tracker server lists
			if cc1.Msg == SzTrackerS2SPrefix {
				abR, err := SvKvDbOp.Get(SzTrackerS2SPrefix)
				// log.Println(err, abR)
				if nil == err && nil != abR && 0 < len(abR) {
					cc1.Conn.WriteTo(abR, cc1.Remote)
				} else {
					aBts := []byte(TrackerServerList[0])
					cc1.Conn.WriteTo(aBts, cc1.Remote)
					SvKvDbOp.Put(SzTrackerS2SPrefix, aBts)
				}
				// for Tracker server 2 Tracker server，防止恶意注册，每个IP每分钟只能注册10次
			} else {
				// 防止恶意注册，每个IP每分钟只能注册10次
				// ip 作为key
				szIp := strings.Split(cc1.Remote.String(), ":")[0]
				abR, err := SvKvDbOp.Get(szIp)
				var nTm int = 1
				if nil == err && nil != abR && 0 < len(abR) {
					aTt := strings.Split(string(abR), SzMacIpSep)
					nTm1, _ := strconv.Atoi(string(aTt[1]))
					intVar, err := strconv.ParseInt(string(abR[0]), 10, 64)
					if nil == err {
						nTm = nTm1
						tm := time.Unix(intVar, 0)
						delta := time.Now().Sub(tm)
						// 小于1分钟就返回
						if delta < time.Minute && nTm >= 10 {
							return
						}
						nTm = nTm + 1
					}
				}
				// 计数器
				SvKvDbOp.Put(szIp, []byte(fmt.Sprintf("%d%s%d", time.Now().UnixNano()/1e6, SzMacIpSep, nTm)))
				aData := strings.Split(cc1.Msg[len(SzTrackerS2SPrefix):], "/")
				if 2 == len(aData) {
					abR, err := SvKvDbOp.Get(SzTrackerS2SPrefix)
					aD1 := []string{}
					if nil != err && nil != abR && 0 < len(abR) {
						aD1 = strings.Split(string(abR), "\n")
						cc1.Conn.WriteTo(abR, cc1.Remote)
					}
					aD1 = append(aD1, strings.Split(aData[1], SzPubSep)...)
					SvKvDbOp.Put(SzTrackerS2SPrefix, []byte(strings.Join(aD1, "\n")))

				}
			}
		}

	}

}
