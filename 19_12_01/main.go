package main

import (
	"bufio"
	"math"
	"os"
	"strconv"

	"github.com/prometheus/common/log"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	log.Info("Advent of Code 01/12/19")
	var totalVal float64

	lineArr, err := readLines("input.txt")
	if err != nil {
		log.Error("ERROR: ", err.Error())
	}

	log.Info(len(lineArr))

	for _, val := range lineArr {
		intermediateFuelVal, _ := strconv.ParseFloat(val, 64)
		for intermediateFuelVal > 0 {
			// take each, divide by three, round down, and subtract 2
			valFloat := intermediateFuelVal
			val1 := valFloat / 3
			val2 := math.Floor(val1)
			val3 := val2 - 2

			if val3 > 0 {
				totalVal = totalVal + val3
			} else {
				log.Info("VALUE OF val3: ", val3)
			}

			// update intermediate value with left over
			intermediateFuelVal = val3
		}
	}

	log.Info("Total Fuel required: ", totalVal)
	log.Info("Complete")

}
