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
func readLines(path string) ([]int64, error) {
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

func main() {

	log.Info("Advent of Code 02/12/19")

	idx_one_master := 0
	idx_two_master := 0
	numberArrMaster, err := readLines("input.txt")
	if err != nil {
		log.Error("ERROR: ", err.Error())
	}

	log.Info(len(numberArrMaster))

	// Create a duplicate of array to be worked on
	var numberArrTemp []int64
	for _, val := range numberArrMaster {
		numberArrTemp = append(numberArrTemp, val)
	}

	// Initial values of the array indices
	done := false
	doneMaster := false
	for i := 0; i <= 99; i++ {
		if doneMaster {
			break
		}
		for j := 0; j <= 99; j++ {
			if doneMaster {
				break
			}
			// Reset current array to original values
			for idx, val := range numberArrMaster {
				numberArrTemp[idx] = val
			}
			idxOp := 0
			idxFirst := 1
			idxSec := 2
			idxPos := 3

			// update the first 2 input fields with the new values of i and j
			numberArrTemp[1] = int64(i)
			numberArrTemp[2] = int64(j)
			log.Info(numberArrTemp[1], ",", numberArrTemp[2])
			done = false
			for done != true {
				operatorVal := numberArrTemp[idxOp]
				firstVal := numberArrTemp[idxFirst]
				secondVal := numberArrTemp[idxSec]
				postionToStore := numberArrTemp[idxPos]
				// log.Info("Operator Val: ", operatorVal)
				// log.Info("First value: ", firstVal)
				// log.Info("Second value: ", secondVal)
				// log.Info("Pos to store output: ", postionToStore)

				var newVal int64
				switch operatorVal {
				case 1:
					log.Info("ADDITION")
					newVal = numberArrTemp[firstVal] + numberArrTemp[secondVal]
					break
				case 2:
					log.Info("MULTIPLICATION")
					newVal = numberArrTemp[firstVal] * numberArrTemp[secondVal]
					break
				case 99:
					log.Info("SEQUENCE COMPLETE")
					done = true
					break
				default:
					log.Error("UNKNOWN Operator,", operatorVal, " Exiting...")
					os.Exit(-1)
				}

				if newVal == 19690720 {
					log.Info("FOUND THE VALUE!")
					log.Info("noun :", numberArrTemp[firstVal])
					log.Info("verb :", numberArrTemp[secondVal])

					if postionToStore == 0 {
						log.Info("correct store location as well, collecting index values...")
						idx_one_master = i
						idx_two_master = j
						doneMaster = true
					}
				}

				if !done {

					// Update array postion with new val
					numberArrTemp[postionToStore] = newVal

					// Move pointer in number array to next instruction and values ready for next iteration
					idxOp = idxOp + 4
					idxFirst = idxFirst + 4
					idxSec = idxSec + 4
					idxPos = idxPos + 4
				}
			}
		}
	}

	log.Info("100 * noun + verb = ", (100*idx_one_master)+idx_two_master)

	log.Info("Complete")
}
