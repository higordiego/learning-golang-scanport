package model

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type Ports struct {
	Port    int
	State   bool
	Service string
}

type Range struct {
	Start int
	End   int
}

type ResultScan struct {
	hostname string
	ip       []net.IP
	results  []Ports
}

var configPorts = map[int]string{
	20: "ftp",
	21: "ftp",
	22: "ssh",
}

func ScanPort(protocol, hostname, service string, port int, resultChannel chan Ports, wg *sync.WaitGroup) {

	defer wg.Done()

	result := Ports{Port: port, Service: service}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))

	conn, err := net.DialTimeout(protocol, address, time.Duration(time.Second*1))

	if err != nil {
		result.State = false
		resultChannel <- result
		return
	}

	defer conn.Close()

	result.State = true
	resultChannel <- result
	return
}

func GetOpenPorts(hostname string, ports Range) {

	scanned, err := ScanPorts(hostname, ports)

	if err != nil {
		log.Println("err", err)
	}

	log.Println(scanned)

}

func ScanPorts(hostname string, ports Range) (ResultScan, error) {

	var results []Ports
	var scanned ResultScan

	var wg sync.WaitGroup

	channelResults := make(chan Ports, ports.End-ports.Start)

	addr, err := net.LookupIP(hostname)

	if err != nil {
		return scanned, err
	}

	for i := ports.Start; i <= ports.End; i++ {
		if service, ok := configPorts[i]; ok {
			wg.Add(1)
			go ScanPorts("tcp", hostname, service, i, channelResults, &wg)
		}
	}

	wg.Wait()

	close(channelResults)

	for result := range channelResults {
		results = append(results, result)
	}

	scanned = ResultScan{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}

	return scanned, nil

}
