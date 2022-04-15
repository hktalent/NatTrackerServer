[![Tweet](https://img.shields.io/twitter/url/http/Hktalent3135773.svg?style=social)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![Follow on Twitter](https://img.shields.io/twitter/follow/Hktalent3135773.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![GitHub Followers](https://img.shields.io/github/followers/hktalent.svg?style=social&label=Follow)](https://github.com/hktalent/)
[![Top Langs](https://profile-counter.glitch.me/hktalent/count.svg)](https://51pwn.com)

# NatTrackerServer
A fast, high performance Cross-platform lightweight Nat Tracker Server

## Tracker Server protocol
### 1、get NAT public ip and port
Multiple requests, returning multiple different ports, can be used for NAT
```
//nat
```
reurn data:
```
122.44.78.33.223:998
```

### 2、Register the currently returned NAT ip:port and return a list of member ips identified by other uuids
format
//51pwn/P2P&E2E/[uuid]/your_publicIpPort_or_0/your_LanIps/self_mac_Addres
eg:
```bash
uuidgen
```
46D0F69E-E347-47A7-8E95-43E5BAA95348
```bash
ifconfig en0
```
```
en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
	options=6463<RXCSUM,TXCSUM,TSO4,TSO6,CHANNEL_IO,PARTIAL_CSUM,ZEROINVERT_CSUM>
	ether 77:5b:67:a5:87:66 
	media: autoselect
```
send data:
```
//51pwn/P2P&E2E/46D0F69E-E347-47A7-8E95-43E5BAA95348/0/192.168.0.56,172.66.0.10,10.10.101.4/77:5b:67:a5:87:66
```
return data:
```
122.44.78.33.223:998;192.168.0.56,172.66.0.10,10.10.101.4 [your self ip and port for other member NAT]
22.144.178.133.23:1998 [with you same uuid member]
122.44.78.33.223:1998;192.168.0.56,72.66.0.101,10.10.101.14 [your Lan other member NAT，and member lan ips]
...
```
- Why do you need Lan ip（s）?
Because multiple nodes are in networks of different depths and levels, they cannot penetrate each other, but they can access the Internet of Things and have the same Internet IP. This design allows them to know the IP of their own intranet, so they can directly connect to each other. network communication without internet

### 3、Register your public ip:port and return a list of member ips identified by other uuids
lan ips is 0，mean，Indicates that there is no need to communicate with other intranets, that is, you do not want to inform other intranet nodes under the same Internet IP address.
send data:
```
//51pwn/P2P&E2E/46D0F69E-E347-47A7-8E95-43E5BAA95348/122.44.78.33.223:998/0/77:5b:67:a5:87:66
```
return data:
```
122.44.78.33.223:998 [your public ip and port for other member NAT]
22.144.178.133.23:1998 [with you same uuid member]
...
```
### 4、between tracker servers protocol
register，format：
//tcksvr/publicIpPort，eg:
```
//tcksvr/223.16.111.99:9980
```
return all Nat tracker lists,like this:
```
223.16.111.99:9980
22.116.11.199:99
```
client,Users in the intranet, get more NAT trackers to speed up mutual discovery between their own nodes
send:
```
//tcksvr
```
return all Nat tracker lists,like this:
```
223.16.111.99:9980
22.116.11.199:99
```

# How build
```bash
build main.go
# run server
./main
```

# How test

```bash
build clientest.go
# run test
./clientest
```

# Donation
| Wechat Pay | AliPay | Paypal | BTC Pay |BCH Pay |
| --- | --- | --- | --- | --- |
|<img src=https://github.com/hktalent/myhktools/blob/master/md/wc.png>|<img width=166 src=https://github.com/hktalent/myhktools/blob/master/md/zfb.png>|[paypal](https://www.paypal.me/pwned2019) **miracletalent@gmail.com**|<img width=166 src=https://github.com/hktalent/myhktools/blob/master/md/BTC.png>|<img width=166 src=https://github.com/hktalent/myhktools/blob/master/md/BCH.jpg>|
