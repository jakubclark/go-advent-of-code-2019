package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Sum := 0.0
	part2Sum := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Not a number, err=", err)
		}
		part1Sum += fuelForMass(mass)
		part2Sum += fuelForMassAndFuel(mass)
	}

	fmt.Println("Solution for Part 1:", strconv.FormatFloat(part1Sum, 'f', 0, 64))
	fmt.Println("Solution for Part 2:", strconv.FormatFloat(part2Sum, 'f', 0, 64))
}

func fuelForMass(mass float64) float64 {
	return math.Floor(mass/3.0) - 2.0
}

func fuelForMassAndFuel(mass float64) float64 {
	fuelMass := fuelForMass(mass)

	if fuelMass > 0.0 {
		return fuelMass + fuelForMassAndFuel(fuelMass)
	} else {
		return 0.0
	}

}
