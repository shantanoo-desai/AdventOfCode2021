package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed diagnostics_report.txt
var diagnosticsFile string

// DiagnosticsReport Structure that generates Epsilon and Gamma Parameters
type DiagnosticsReport struct {
	Epsilon string
	Gamma   string
	Oxygen  string
	CO2     string
}

// GenerateParameters method derives the Gamma, Epsilon Parameters as a binary string
func (d *DiagnosticsReport) GenerateParameters(fileName string) {
	binaries := strings.Split(fileName, "\n") // make binary rows

	for i := 0; i < len(binaries[0]); i++ { // iterate over the complete row

		var ones, zeroes int
		for _, b := range binaries { // check count of 1 / 0 on the particular column
			if b[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones { // Gamma is the most occurred digit, Epsilon is the least occured digit
			d.Gamma += "0"
			d.Epsilon += "1"
		} else {
			d.Gamma += "1"
			d.Epsilon += "0"
		}
	}
}

// Product method returns the product of epsilon and gamma
func (d DiagnosticsReport) Product() int {
	epsilon, err := strconv.ParseInt(d.Epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	gamma, err := strconv.ParseInt(d.Gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(epsilon * gamma)
}

// GenerateLifeSupportParameters extracts the Oxygen, CO2 Values in binary value as strings
func (d *DiagnosticsReport) GenerateLifeSupportParameters(fileName string) {
	binarySets := strings.Split(fileName, "\n")

	for i := 0; len(binarySets) > 1; i++ { // start with all given values
		// calculate most common bits for oxgen value
		var ones, zeroes int
		decidingChar := "1"
		var newOSet []string
		for _, b := range binarySets {
			if b[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		// Criteria: most common value (0 or 1) in bit position
		// pick values only starting with most common value
		// in case of equal 0 / 1 => pick 1 for oxygen
		if zeroes > ones {
			decidingChar = "0"
		}

		for _, n := range binarySets { // create a subset of the values
			if string(n[i]) == decidingChar {
				newOSet = append(newOSet, n)
			}
		}
		binarySets = newOSet
	}
	d.Oxygen = binarySets[0]

	binarySets = strings.Split(fileName, "\n")
	for i := 0; len(binarySets) > 1; i++ { // start with all given values
		// calculate most common bits for oxgen value
		var ones, zeroes int
		decidingChar := "0"
		var newOSet []string
		for _, b := range binarySets {
			if b[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		// Criteria: least common value (0 or 1) in bit position
		// pick values only starting with least common value
		// in case of equal 0 / 1 => pick 0 for CO2
		if ones < zeroes {
			decidingChar = "1"
		}

		for _, n := range binarySets { // create a subset of the values
			if string(n[i]) == decidingChar {
				newOSet = append(newOSet, n)
			}
		}
		binarySets = newOSet
	}
	d.CO2 = binarySets[0]
}

// LifeSupportProduct returns the product of Oxgen and CO2
func (d DiagnosticsReport) LifeSupportProduct() int {
	o, err := strconv.ParseInt(d.Oxygen, 2, 64)
	if err != nil {
		panic(err)
	}
	co2, err := strconv.ParseInt(d.CO2, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(o * co2)
}

func main() {

	report := DiagnosticsReport{}
	fmt.Println("Generating Submarine Diagnostics Report")
	report.GenerateParameters(diagnosticsFile)

	fmt.Printf("Diagnostics Report: Epsilon, Gamma: %s, %s\n", report.Epsilon, report.Gamma)
	fmt.Printf("Diagnostics Report Product: %d\n", report.Product())

	fmt.Println("Generating Life Support System Parameters")
	report.GenerateLifeSupportParameters(diagnosticsFile)

	fmt.Printf("Life Support Report: 0: %s, CO2: %s\n", report.Oxygen, report.CO2)
	fmt.Printf("Life Support Product: %d\n", report.LifeSupportProduct())
}
