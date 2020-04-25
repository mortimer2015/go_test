package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
	"log"
	"os"
	"sync"
	"time"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	fmt.Println(time.Now())
	ipList, ipDictData := ipDataHandle.IpDict("file/ipdata.txt")
	i := 2
	if i == 1 {
		ret := ipDataHandle.IpFrom(ipList, ipDictData, "1.0.1.222")
		fmt.Printf(ret)
	}

	//fileName := "file/test.csv"
	//f, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//data := strings.Split(string(f), "\n")
	//for i := 0; i < len(data); i++ {
	//	//fmt.Println(data[i])
	//	line := strings.Split(data[i], ", ")
	//	//fmt.Println(line[len(line)-1])
	//	ret := ipDataHandle.IpFrom(ipList, ipDictData, line[3])
	//	fmt.Printf("%s, %s", data[i], ret)
	//}
	fmt.Println(time.Now())
}

func handleWorker() {
	defer wg.Done()

	mutex.Lock()
	{
		f, err := os.OpenFile("file/ret1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(retString)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
	mutex.Unlock()
}
