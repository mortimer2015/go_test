package ipDataHandle

import "sort"

func IpFrom(ipList []int, ipDictData []IpData, ip string) string {
	ipInt, _ := Ip2int(ip)
	dataLen := len(ipList)
	i := sort.Search(dataLen, func(i int) bool { return ipList[i] >= ipInt })
	if i == dataLen {
		return "unknown"
	}
	if ipDictData[i-1].Start <= ipInt && ipInt <= ipDictData[i-1].Stop {
		return ipDictData[i-1].Describe
	} else {
		return "unknown"
	}

}
