package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
)

func main() {
	fmt.Println("aaa")
	ipData := "127.0.0a.1"
	ipInt, err := ipDataHandle.Ip2int(ipData)
	if err != nil {
		fmt.Printf("%s  %s", ipData, err.Error())
	} else {
		fmt.Println(ipInt)
	}

}
