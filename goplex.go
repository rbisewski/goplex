/*
 * Goplex Functions Tests
 *
 * Description: A set of tests that use the functions defined in other
 *              files of this package.
 *
 * Author: Robert Bisewski <contact@ibiscybernetics.com>
 */

//
// Package
//
package main

//
// Imports
//
import (
	"fmt"
	"math"
	"os"
)

//
// PROGRAM MAIN
//
func main() {

	// variable declaration:
	//
	// a) expected --> correct value
	// b) actual   --> experimental value calculated by this program
	//
	var actual float64
	var expected float64

	// tell the end-user the tests are starting
	fmt.Println("Goplex tests begin now...")

	//
	// Tsiolkovsky Delta-V Launch
	//
	Ve := 17000.0
	m0 := 5000.0
	mf := 3000.0
	expected = 8684.035604021843
	actual = tsiolkovskyDeltaV(Ve, m0, mf)

	// test to ensure this got the expected result
	if expected != actual {
		fmt.Println("Tsiolkovsky Delta-V test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Energy of a photon with a wavelength of 400nm
	//
	wavelength := 400.0 * math.Pow(10, -9)
	expected = 4.966114480984394 * math.Pow(10, -19)
	actual = photonEnergy(wavelength)

	// test to ensure this got the expected result
	if expected != actual {
		fmt.Println("Photon Energy test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Thermal velocity of gas propellant
	//
	g := 9.8000
	m := 0.8100                // RP-1 density
	tempInKelvins := 3670.0000 // RP-1 temp
	expected = 0.11043619735553113
	actual = thermalVelocityOfHeatedGas(g, tempInKelvins, m)

	// test to ensure this got the expected result
	if expected != actual {
		fmt.Println("Thermal velocity test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Calculate the Lorentz factor of 0.5c
	//
	halfC := c / 2.0
	expected = 1.1547005383792517
	actual = lorentzFactor(halfC)

	// test to ensure this got the expected result
	if expected != actual {
		fmt.Println("Lorentz factor test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Abraham-Lorentz force for a up quark in a vacuum
	//
	upQuarkQ := 0.66666666
	gravJerk := 9.8
	expected = 9.6857127934588849 * math.Pow(10, -16)
	actual = abrahamLorentzForce(upQuarkQ, vacuumPermittivity, gravJerk)

	// test to ensure this got the expected result
	if expected != actual {
		fmt.Println("Abraham-Lorentz test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Perihelion Shift Calculation, of Mercury
	//
	L := 57909050.0
	T := 47.362
	e := 0.205630
	expected = 0.065884179454766778
	actual = perihelionShift(L, T, e) * secondsInADay * 59.0

	// do a quick check to ensure this actually gave the right result
	if expected != actual {
		fmt.Println("Perihelion Shift test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	//
	// Schwarzschild radius of the planet Earth
	//
	expected = 0.008870062974351377
	actual = schwarzschildRadius(massOfTheEarth)

	// do a quick check to ensure this actually gave the right result
	if expected != actual {
		fmt.Println("Schwarzshild radius test failed!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Calculated: ", actual)
		os.Exit(1)
	}

	// otherwise everything turned out fine
	fmt.Println("All tests completed successfully!")
}
