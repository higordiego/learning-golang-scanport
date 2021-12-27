package model

import (
	"fmt"
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
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "SSH",
	23:   "telnet",
	25:   "SMTP",
	43:   "whois",
	53:   "dns",
	66:   "Oracle SQL*NET?",
	67:   "dhcp",
	68:   "dhcp",
	69:   "tftp",
	80:   "http",
	88:   "kerberos",
	109:  "pop2",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	445:  "Samba",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	631:  "cups",
	1194: "openvpn",
	5000: "unpn",
	5800: "VNC remote desktop",
	194:  "IRC",
	118:  "SQL service?",
	150:  "SQL-net?",
	1433: "Microsoft SQL server",
	1434: "Microsoft SQL monitor",
	3306: "MySQL",
	3396: "Novell NDPS Printer Agent",
	3535: "SMTP (alternate)",
	8080: "https-proxy",
	8443: "https-alt",
	9160: "Cassandra [ http://cassandra.apache.org/ ]",
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
	} else {
		DisplayScanResult(scanned)
	}
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
			go ScanPort("tcp", hostname, service, i, channelResults, &wg)
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

func DisplayScanResult(result ResultScan) {
	ip := result.ip[len(result.ip)-1]
	fmt.Printf("Open ports for %s (%s)\n\n", result.hostname, ip.String())
	for _, v := range result.results {
		if v.State {
			fmt.Printf("%d	%s\n", v.Port, v.Service)
		}
	}
	fmt.Printf("\n")
}
