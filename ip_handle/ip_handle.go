package main

import (
	"fmt"
	"go_test/ip_handle/ipDataHandle"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	mutex      sync.Mutex
	ipList     []int
	ipDictData []ipDataHandle.IpData
)

func main() {
	startTime := time.Now()
	fmt.Println(startTime)
	ipList, ipDictData = ipDataHandle.IpDict("file/ipdata.txt")

	taskData := make(chan string, 10000)

	n := 2048
	wg.Add(n)
	for i := 0; i < n; i++ {
		go handleWorker(taskData)
	}

	generateBuff("file/test.csv", taskData)
	close(taskData)

	wg.Wait()
	stopTime := time.Now()
	fmt.Println(stopTime)
	fmt.Println(stopTime.Sub(startTime))
}

func generateBuff(fileName string, taskData chan string) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(f), "\n")
	for i := 1; i < len(data); i++ {
		taskData <- data[i]
	}

}

func handleWorker(taskData chan string) {
	defer wg.Done()
	retString := ""
	//ret := ipDataHandle.IpFrom(ipList, ipDictData, "1.0.1.222")
	//fmt.Printf(ret)
	for {
		data, ok := <-taskData
		if !ok {
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
			return
		}

		line := strings.Split(data, ", ")
		//fmt.Println(line[len(line)-1])
		ret := ipDataHandle.IpFrom(ipList, ipDictData, line[3])
		fmt.Printf("%s, %s\n", data, ret)
		retString = retString + data + ", " + ret + "\n"
	}
}
