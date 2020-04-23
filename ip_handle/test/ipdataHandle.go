package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
)

func main() {
	data1, data := ipDataHandle.IpDict("file/ipdata.txt")
	fmt.Println(data1)
	fmt.Println(data)
}
