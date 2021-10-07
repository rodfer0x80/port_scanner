package port

import (
	"fmt"
	"os"
)

func partitionResults(results []ScanResult) []ScanResult {
	var open []ScanResult

	for _, result := range results {
		if result.State == "Open" {
			open = append(open, result)
		}
	}

	return open
}

func buildOutputString(open []ScanResult) string {
	str := "==== UDP ====\n"

	var cached []int
	for i, result := range open {
		if result.Protocol == "UDP" {
			str += result.Port + "\n"
		} else {
			cached = append(cached, i)
		}
	}

	str += "==== TCP ====\n"
	for _, i := range cached {
		str += open[i].Port + "\n"
	}

	return str
}

func getScanOutput(results []ScanResult) string {
	open := partitionResults(results)

	outputString := buildOutputString(open)

	return outputString
}

func StdoutToFile(results []ScanResult, filename string) {
	output := getScanOutput(results)

	if err := os.WriteFile(filename, []byte(output), 0666); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Saved output to file")
}

func StdoutToDisplay(results []ScanResult) {
	output := getScanOutput(results)

	fmt.Println(output)
}
