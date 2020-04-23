package ip2int

import (
	"errors"
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
		intSlice = intSlice << (8 * i)
		ret = ret | intSlice
	}
	return ret, nil
}
