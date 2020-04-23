package ipDataHandle

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type IpData struct {
	Start    int
	Stop     int
	Describe string
}

func IpDict(fileName string) ([]int, []IpData) {
	var ret []IpData
	var ipList []int
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(f), "\n")
	for i := 0; i < len(data); i++ {
		//fmt.Println(data[i])
		line := strings.Split(data[i], " ")
		var lineNew []string
		for y := 0; y < len(line); y++ {
			if line[y] != "" {
				lineNew = append(lineNew, line[y])
			}
		}
		if len(lineNew) > 2 {
			ipStart := ip2intNoError(lineNew[0])
			one := IpData{
				Start:    ipStart,
				Stop:     ip2intNoError(lineNew[1]),
				Describe: strings.Join(lineNew[2:], " "),
			}
			ret = append(ret, one)
			ipList = append(ipList, ipStart)
		}
	}
	return ipList, ret
}

func ip2intNoError(ipData string) int {
	ipInt, err := Ip2int(ipData)
	if err != nil {
		fmt.Printf("%s  %s", ipData, err.Error())
	}
	return ipInt
}
