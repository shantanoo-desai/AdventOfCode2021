package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// DepthReport returns the total number of measurements
// that are larger than the previous measurement
func DepthReport(sonarReadings []int) int {
	if len(sonarReadings) < 2 { // just one reading won't do
		return 0
	}
	previousReading := sonarReadings[1]
	measurementCounter := 1

	for i := 2; i < len(sonarReadings); i++ {
		if sonarReadings[i] > previousReading { // if current reading is bigger than previous value
			measurementCounter++
		}
		previousReading = sonarReadings[i]
	}

	return measurementCounter
}

// SlidingWindow returns the total number of sums larger than the previous sum
// of a three-measurement sliding window
//
// Logical Solution: for i/p [0,1,2,3,4,5,6,7,8,9] (len=10)
// Slices are:
//             [iterator:iterator+windowSize] => values
//            (1) [0:3] => 0,1,2 (Initial Window)
//            (2) [1:4] => 1,2,3 (Iterate from here)
//            (3) [2:5] => 2,3,4
//            (4) [3:6] => 3,4,5
//            (5) [4:7] => 4,5,6
//            (6) [5:8] => 5,6,7
//            (7) [6:9] => 6,7,8
//            (8) [7:10] =>7,8,9
func SlidingWindow(sonarReadings []int) int {
	windowSize := 3
	// Initial Window always starts at the beginning
	previousWindowSum := threeSWSum(sonarReadings[:3])
	measurementCounter := 1

	for i := 1; i < len(sonarReadings)-windowSize; i++ {
		currentWindowSum := threeSWSum(sonarReadings[i : i+windowSize])
		if currentWindowSum > previousWindowSum {
			// current window's sum is greater than previous
			measurementCounter++
		}
		previousWindowSum = currentWindowSum
	}
	return measurementCounter
}

// threeSWSum returns the sum of the incoming three-measurement Sliding Window
func threeSWSum(currentWindow []int) (sum int) {
	for _, val := range currentWindow {
		sum += val
	}
	return
}

func main() {
	var incomingMeasurements []int

	fmt.Println("Reading Sonar Measurements")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error Opening File...")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		me, _ := strconv.Atoi(scanner.Text())
		incomingMeasurements = append(incomingMeasurements, me)
	}

	fmt.Println("Depth Report....")
	result := DepthReport(incomingMeasurements)

	fmt.Printf("Result: %d\n", result)

	fmt.Println("Sliding Window Report...")
	result = SlidingWindow(incomingMeasurements)
	fmt.Printf("Result: %d\n", result)
}
