package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
)

func main() {
	data := ipDataHandle.IpDataHandle("file/ipdata.txt")
	fmt.Println(data)
}
