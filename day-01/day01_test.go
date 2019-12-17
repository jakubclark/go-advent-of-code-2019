package main

import (
	"testing"
)

func Test1(t *testing.T) {
	if !runTestPart1(12.0, 2.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test2(t *testing.T) {
	if !runTestPart1(14.0, 2.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test3(t *testing.T) {
	if !runTestPart1(1969.0, 654.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test4(t *testing.T) {
	if !runTestPart1(100756.0, 33583.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test5(t *testing.T) {
	if !runTestPart2(14.0, 2.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test6(t *testing.T) {
	if !runTestPart2(1969.0, 966.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func Test7(t *testing.T) {
	if !runTestPart2(100756.0, 50346.0) {
		t.Errorf("Got unexpected mass for input=12.0")
	}
}

func runTestPart1(startingMass float64, expectedMass float64) bool {
	mass := fuelForMass(startingMass)
	return mass == expectedMass
}

func runTestPart2(startingMass float64, expectedMass float64) bool {
	mass := fuelForMassAndFuel(startingMass)
	return mass == expectedMass
}
