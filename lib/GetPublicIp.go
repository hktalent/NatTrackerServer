package lib

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	Nat string = "//nat"
	// P2P or E2E 前缀
	NatPrefix     string = "//51pwn/P2P&E2E/"
	NatBodyFormat string = "%s/%s/%s"

	// Nat tracker server之间通讯协议前缀
	S2sPrefix string = "//tcksvr"

	// 默认值
	SzDefault string = "0"
	// port 分隔符号
	SzPortSep string = ":"
	// public Ip分隔
	SzPubSep string = ";"
	// Mac-ip 分隔
	SzMacIpSep string = "-"
	// lan ip 分隔
	SzLanIpSep string = ","
	// 本地ip
	privateIPBlocks []*net.IPNet
	doOnce          sync.Once
)

type NatTrackerProtocol struct {
	Msg    string
	Conn   *net.UDPConn
	Remote *net.UDPAddr
}

func Init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

// new
func NewNatTrackerProtocol() *NatTrackerProtocol {
	doOnce.Do(func() {
		Init()
	})
	return &NatTrackerProtocol{}
}

func (r *NatTrackerProtocol) isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// //
// func (r *NatTrackerProtocol) GetNatIpAndPortStr() string {
// 	return Nat
// }

// generate a UUID
func (r *NatTrackerProtocol) GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

// []string indexof
func (r *NatTrackerProtocol) ArrStrIndexOf(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return -1
}

// get local public Ip to line: ip1:port;ip2:port;ip3:port
func (r *NatTrackerProtocol) GetPublicIP2Line(szPort string) string {
	a, err := r.GetPublicIP()
	if nil == err {
		return SzDefault
	}
	return strings.TrimSuffix(strings.Join(a, SzPortSep+szPort+SzPubSep), SzPubSep)
}

// get all public IP lists
func (r *NatTrackerProtocol) GetPublicIP() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string

	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		// not have mac
		if a != "" {
			addrs, err := ifa.Addrs()
			// get Ip error
			if nil != err {
				continue
			}
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					if r.isPrivateIP(v.IP) {
						continue
					}
					as = append(as, v.IP.String())
				case *net.IPAddr:
					if r.isPrivateIP(v.IP) {
						continue
					}
					as = append(as, v.IP.String())
				}
			}
		}
	}
	return as, nil
}

// get Mac-Ip to line
func (r *NatTrackerProtocol) GetMacIPAddrLists2Line() string {
	a, err := r.GetLocalMacIPAddrLists()
	if nil == err {
		return strings.Join(a, SzPubSep)
	}
	return SzDefault
}

// Get MAC-ip1,ip2;ip4 address
func (r *NatTrackerProtocol) GetLocalMacIPAddrLists() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	aFlt := []string{"192.168.", "10.", "172.16."}
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		// not have mac
		if a != "" {
			addrs, err := ifa.Addrs()
			// get Ip error
			if nil != err {
				continue
			}
			var aR []string
			var x1 string
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					x1 = v.IP.String()
					if r.isPrivateIP(v.IP) {
						if strings.HasSuffix(x1, aFlt[0]) || strings.HasSuffix(x1, aFlt[1]) || strings.HasSuffix(x1, aFlt[2]) {
							aR = append(aR, x1)
						}
					} else {
						aR = append(aR, x1)
					}
				case *net.IPAddr:
					x1 = v.IP.String()
					if r.isPrivateIP(v.IP) {
						if strings.HasSuffix(x1, aFlt[0]) || strings.HasSuffix(x1, aFlt[1]) || strings.HasSuffix(x1, aFlt[2]) {
							aR = append(aR, x1)
						}
					} else {
						aR = append(aR, x1)
					}
				}
			}
			x1 = strings.Join(aR, SzLanIpSep)
			if "" != x1 {
				x1 = a + SzMacIpSep + x1
				if -1 == r.ArrStrIndexOf(as, x1) {
					as = append(as, x1)
				}

			}
		}
	}
	return as, nil
}

// suport IPv4 and IPv6
func (r *NatTrackerProtocol) GetNatPublicIpAndPort(szServerIp string) (string, error) {

	address := szServerIp
	conn, err := net.Dial("udp", address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(Nat))
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
