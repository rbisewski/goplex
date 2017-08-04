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

    // seconds in a single Earth day
    seconds_in_a_day float64 = 86400

    // Speed of light in a vaccuum, in m/s
    c float64 = 299792458.0

    // Universal gravitational constant, in m^3 kg^-1 s^-2
    universal_gravitation_constant float64 = 0.0000000000667408
)

//! Tsiolkovsky rocket equation
/*
 * @param    float64    effective exhaust velocity         --> Ve
 * @param    float64    initial total mass (w/ propellant) --> m0
 * @param    float64    final total mass (w/o propellant)  --> mf
 *
 * @result   float64    delta-v
 */
func tsiolkovsky_delta_v(Ve float64, m0 float64, mf float64) (float64) {

    // input validation
    if mf == 0 {
        return 0;
    }

    // calculate the mass ratio, i.e. the different between the initial
    // propellant and the final mass w/o propellant
    var ratio_of_initial_to_dry_mass float64 = m0 / mf

    // determine the natural log of the mass ratio
    var nlog_of_mass_ratio float64 = math.Log(ratio_of_initial_to_dry_mass)

    // figure out the total energy requiring during the change of mass
    // over the start and end of the launch, otherwise known as the delta-V
    var delta_v float64 = Ve * nlog_of_mass_ratio

    // go ahead and return the values
    return delta_v
}

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

    // speed of light in kilometers per second
    var c_in_km_per_second float64 = c * 1000

    // calculate the spherical shape of the semi-major axis
    var dividend float64 = 24 * math.Pi * math.Pi * math.Pi * L * L

    // calculate the velocity of the orbital body
    var divisor float64 = T * T * c_in_km_per_second * c_in_km_per_second * (1 - e * e)

    // safety check, if the divisor is zero, return 0
    if divisor == 0.0 {
        return 0.0;
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
    // Tsiolkovsky Delta-V Launch
    //
    var delta_v_expected_result float64 = 8684.035604021843
    var Ve float64 = 17000.0
    var m0 float64 = 5000.0
    var mf float64 = 3000.0
    var delta_v_test = tsiolkovsky_delta_v(Ve, m0, mf)

    // test to ensure this got the expected result
    if delta_v_expected_result != delta_v_test {
        fmt.Println("Tsiolkovsky Delta-V test failed!")
        fmt.Println("Expected: ", delta_v_expected_result)
        fmt.Println("Calculated: ", delta_v_test)
        os.Exit(1)
    }

    //
    // Perihelion Shift Calculation, of Mercury
    //
    var perihelion_shift_expected_result float64 = 0.0
    var L float64 = 57909050.0
    var T float64 = 47.362
    var e float64 = 0.205630
    var perihelion_shift_test float64 = perihelion_shift(L, T, e)

    // do a quick check to ensure this actually gave the right result
    if perihelion_shift_test != perihelion_shift_expected_result {
        fmt.Println("Perihelion Shift test failed!")
        fmt.Println("Expected: ", perihelion_shift_expected_result)
        fmt.Println("Calculated: ", perihelion_shift_test *
        seconds_in_a_day * 59)
        os.Exit(1)
    }

    //
    // TODO: insert further tests here
    //
}
