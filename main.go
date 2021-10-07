package main

import (
	"os"

	"github.com/trevalkov/port_scanner/port"
)

func main() {
	hostname := os.Args[1]

	filename := "display"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}

	if filename == "display" {
		port.StdoutToDisplay(hostname)
	} else {
		port.StdoutToFile(hostname, filename)
	}

}
