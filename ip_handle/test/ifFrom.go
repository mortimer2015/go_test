package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
)

var (
	ipList     []int
	ipDictData []ipDataHandle.IpData
)

func main() {
	ipList, ipDictData = ipDataHandle.IpDict("file/ipdata.txt")
	ret := ipDataHandle.IpFrom(ipList, ipDictData, "1.0.1.222")
	fmt.Println(ret)
}

//func IpFrom(ip string) string {
//	ipInt, _ := ipDataHandle.Ip2int(ip)
//	dataLen := len(ipList)
//	i := sort.Search(dataLen, func(i int) bool { return ipList[i] >= ipInt })
//	if i == dataLen{
//		return "unknown"
//	}
//	if (ipDictData[i-1].Start <= ipInt && ipInt <= ipDictData[i-1].Stop) {
//		return ipDictData[i-1].Describe
//	} else {
//		return "unknown"
//	}
//
//}
