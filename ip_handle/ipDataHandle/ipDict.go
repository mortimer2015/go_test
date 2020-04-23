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

func IpDataHandle(fileName string) []IpData {
	var ret []IpData
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
			one := IpData{
				Start:    ip2intNoError(lineNew[0]),
				Stop:     ip2intNoError(lineNew[1]),
				Describe: strings.Join(lineNew[2:], " "),
			}
			ret = append(ret, one)
		}
	}
	return ret
}

func ip2intNoError(ipData string) int {
	ipInt, err := Ip2int(ipData)
	if err != nil {
		fmt.Printf("%s  %s", ipData, err.Error())
	}
	return ipInt
}
