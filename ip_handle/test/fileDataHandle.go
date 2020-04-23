package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
	"io/ioutil"
	"strings"
)

func main() {
	ipList, ipDictData := ipDataHandle.IpDict("file/ipdata.txt")
	//ret := ipDataHandle.IpFrom(ipList, ipDictData, "1.0.1.222")
	fileName := "file/test.csv"
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(f), "\n")
	for i := 0; i < len(data); i++ {
		//fmt.Println(data[i])
		line := strings.Split(data[i], ", ")
		//fmt.Println(line[len(line)-1])
		ret := ipDataHandle.IpFrom(ipList, ipDictData, line[3])
		fmt.Printf("%s, %s\n", data[i], ret)
	}
}
