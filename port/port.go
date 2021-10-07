package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, intPort int) ScanResult {
	port := strconv.Itoa(intPort)
	result := ScanResult{Port: port}

	if protocol == "tcp" {
		result.Protocol = "TCP"
	} else {
		result.Protocol = "UDP"
	}

	address := hostname + ":" + port
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()

	result.State = "Open"
	return result
}

func QuickScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1023; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}
