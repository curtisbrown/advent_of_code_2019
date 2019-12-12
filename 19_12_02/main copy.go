package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/prometheus/common/log"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines1(path string) ([]int64, error) {
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

	var numbers []int64
	if len(lines) == 1 {
		numberStr := strings.Split(lines[0], ",")
		for _, num := range numberStr {
			numInt, _ := strconv.ParseInt(num, 10, 64)
			numbers = append(numbers, numInt)
		}
	}
	return numbers, scanner.Err()
}

func mainPart1() {

	log.Info("Advent of Code 02/12/19")

	numberArr, err := readLines1("input.txt")
	if err != nil {
		log.Error("ERROR: ", err.Error())
	}

	log.Info(len(numberArr))

	// Initial values of the array indices
	idxOp := 0
	idxFirst := 1
	idxSec := 2
	idxPos := 3
	done := false
	for done != true {
		operatorVal := numberArr[idxOp]
		firstVal := numberArr[idxFirst]
		secondVal := numberArr[idxSec]
		postionToStore := numberArr[idxPos]
		log.Info("Operator Val: ", operatorVal)
		log.Info("First value: ", firstVal)
		log.Info("Second value: ", secondVal)
		log.Info("Pos to store output: ", postionToStore)

		var newVal int64
		switch operatorVal {
		case 1:
			log.Info("ADDITION")
			newVal = numberArr[firstVal] + numberArr[secondVal]
			break
		case 2:
			log.Info("MULTIPLICATION")
			newVal = numberArr[firstVal] * numberArr[secondVal]
			break
		case 99:
			log.Info("SEQUENCE COMPLETE")
			done = true
			break
		default:
			log.Error("UNKNOWN Operator,", operatorVal, "Exiting...")
			os.Exit(-1)
		}

		if !done {

			// Update array postion with new val
			numberArr[postionToStore] = newVal

			// Move pointer in number array to next instruction and values ready for next iteration
			idxOp = idxOp + 4
			idxFirst = idxFirst + 4
			idxSec = idxSec + 4
			idxPos = idxPos + 4
		}
	}

	log.Info("Complete")

}
