[![Tweet](https://img.shields.io/twitter/url/http/Hktalent3135773.svg?style=social)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![Follow on Twitter](https://img.shields.io/twitter/follow/Hktalent3135773.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![GitHub Followers](https://img.shields.io/github/followers/hktalent.svg?style=social&label=Follow)](https://github.com/hktalent/)
[![Top Langs](https://profile-counter.glitch.me/hktalent/count.svg)](https://51pwn.com)

# NatTrackerServer
- A fast, high performance Cross-platform lightweight Nat Tracker Server
- suport IPv4 and IPv6
<img width="786" alt="image" src="https://user-images.githubusercontent.com/18223385/163664762-b418b6da-735f-43e4-a948-f2034491628b.png">

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

### 2、Register the (currently returned NAT ip:port) or (public ip:port) and return a list of member ips identified by other uuids
format:
//51pwn/P2P&E2E/[uuid]/your_publicIpPort_or_0/your_LanMac-Ips_or_0

send data:
```
//51pwn/P2P&E2E/46D0F69E-E347-47A7-8E95-43E5BAA95348/33.231.55.111:883,133.23.155.121:1883;33.123.155.11:4883/b6:11:48:10:91:82-1e88::aede:48ff:fe00:1122;90:9c:4a:ce:39:72-fe80::894:75ca:61c3:fbec,192.168.10.17;ca:76:ae:e7:31:9c-fe99::6676:eeff:f117:009c;42:0c:aa:cc:ee:68-172.26.53.10
```
return data:
```
122.44.78.33.223:998;192.168.0.56,172.66.0.10,10.10.101.4 [your self ip and port for other member NAT]
22.144.178.133.23:1998 [with you same uuid member]
122.44.78.33.223:1998;192.168.0.56,72.66.0.101,10.10.101.14 [your Lan other member NAT，and member lan ips]
...
```
- Why do you need Lan ip（s）?
Because multiple nodes are in networks of different depths and levels, they cannot penetrate each other, 
but they can access the Internet of Things and have the same Internet IP. 
This design allows them to know the IP of their own intranet, 
so they can directly connect to each other. network communication without internet

- why have reg sefl public ip and port?
If you have already monitored the port and have a public Internet IP,
maybe you need to directly expose the IP and port to the tracker,
so that the first one returned by the tracker is no longer Nat's address, but the address given by yourself.

### 4、between tracker servers protocol
Prevent malicious registration, each IP can only register 10 times per minute
register，format：
//tcksvr/publicIpPort，eg:
```
//tcksvr/223.16.111.99:9980;123.116.111.99:980
```
return all Nat tracker server lists,like this:
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
# How run
## server


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
