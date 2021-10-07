package port

import (
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, intPort int) ScanResult {
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

func quickScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1023; i++ {
		results = append(results, scanPort("tcp", hostname, i))
		results = append(results, scanPort("udp", hostname, i))
	}

	return results
}
