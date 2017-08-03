/*
 * Goplex Math Functions 
 *
 * Description: A set of useful functions for certain scientific and
 *              mathematical fields.
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
// Globals
//
var (

    // Speed of light in a vaccuum, in m/s
    c float64 = 299792458.0

    // Universal gravitational constant, in m^3 kg^-1 s^-2
    universal_gravitation_constant float64 = 0.0000000000667408
)

//! Function to calculate the perihelion shift of an orbit
/*
 * @param    float64    semi-major axis      --> L
 * @param    float64    orbital speed        --> T
 * @param    float64    orbital eccentricity --> e 
 *
 * @result   float64    perihelion shift, in radians/revolution
 */
func perihelion_shift(L float64, T float64, e float64) (float64) {

    // variable declaration
    var result float64 = 0.0

    // calculate the spherical shape of the semi-major axis
    dividend := 24 * math.Pi * math.Pi * math.Pi * L * L

    // calculate the velocity of the orbital body
    divisor := T * T * c * c * (1 - e * e)

    // safety check, if the divisor is zero, return 0
    if divisor == 0 {
        return 0;
    }

    // combine both to obtain the perihelion
    result = dividend / divisor

    // if we got this far, go ahead and return the result
    return result;
}

//
// PROGRAM MAIN
//
func main() {

    // tell the end-user the tests are starting
    fmt.Println("Goplex tests begin now...");

    //
    // Perihelion Shift Calculation, of Mercury
    //
    var perihelion_shift_expected_result float64 = 0
    var L float64 = 57909050.0
    var T float64 = 47.362
    var e float64 = 0.205630
    var perihelion_shift_test float64 = perihelion_shift(L, T, e)

    // do a quick check to ensure this actually gave the right result
    if perihelion_shift_test != perihelion_shift_expected_result {
        fmt.Println("Perihelion Shift test failed!")
        fmt.Println("Expected: ", perihelion_shift_expected_result)
        fmt.Println("Calculated: ", perihelion_shift_test)
        os.Exit(1)
    }

    //
    // TODO: insert further tests here
    //
}
