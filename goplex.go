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

    // TODO: insert code here

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
    // Perihelion Shift Calculation
    //
    var perihelion_shift_expected_result = 0
    var L float64 = 0
    var T float64 = 0
    var e float64 = 0
    var perihelion_shift_test float64 = perihelion_shift(L, T, e)

    // do a quick check to ensure this actually gave the right result
    if perihelion_shift_test != perihelion_shift_expected_result {
        fmt.Println("Perihelion Shift test failed!")
        os.Exit(1)
    }

    //
    // TODO: insert further tests here
    //
}
