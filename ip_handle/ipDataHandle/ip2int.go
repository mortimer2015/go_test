package ipDataHandle

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Ip2int(ip string) (int, error) {
	ipData := strings.Split(ip, ".")
	if len(ipData) > 4 {
		return 1, errors.New("ip数据错误")
	}
	ret := 0
	for i := 0; i < len(ipData); i++ {
		intSlice, err := strconv.Atoi(ipData[i])
		if err != nil {
			return ret, errors.New("ip数据格式错误")
		}
		ret = ret << 8
		ret = ret + intSlice
	}
	return ret, nil
}

func ip2intNoError(ipData string) int {
	ipInt, err := Ip2int(ipData)
	if err != nil {
		fmt.Printf("%s  %s", ipData, err.Error())
	}
	return ipInt
}
