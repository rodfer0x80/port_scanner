package main

import (
	"github.com/trevalkov/port_scanner/port"
)

func main() {
	filename := "scan_results.txt"

	results := port.QuickScan("localhost")
	port.StdoutToFile(results, filename)
}
