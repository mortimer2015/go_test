package main

import (
	"fmt"
	"go_test/ip_handle/ip2int"
)

func main() {
	fmt.Println("aaa")
	ipData := "127.0.0a.1"
	ipInt, err := ip2int.Ip2int(ipData)
	if err != nil {
		fmt.Printf("%s  %s", ipData, err.Error())
	} else {
		fmt.Println(ipInt)
	}

}
