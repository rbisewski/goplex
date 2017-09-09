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
    var actual float64 = 0
    var expected float64 = 0

    // tell the end-user the tests are starting
    fmt.Println("Goplex tests begin now...");

    //
    // Tsiolkovsky Delta-V Launch
    //
    var Ve float64 = 17000.0
    var m0 float64 = 5000.0
    var mf float64 = 3000.0
    expected = 8684.035604021843
    actual = tsiolkovsky_delta_v(Ve, m0, mf)

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
    var wavelength float64 = 400 * math.Pow(10, -9)
    expected = 4.966114480984394 * math.Pow(10, -19)
    actual = photon_energy(wavelength)

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
    var g float64               = 9.8000
    var m float64               = 0.8100     // RP-1 density
    var temp_in_kelvins float64 = 3670.0000  // RP-1 temp
    expected = 0.11043619735553113
    actual = thermal_velocity_of_heated_gas(g, temp_in_kelvins, m)

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
    var half_c float64 = c / 2
    expected = 1.1547005383792517
    actual = lorentz_factor(half_c)

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
    var up_quark_q float64 = 0.66666666
    var grav_jerk float64 = 9.8
    expected = 9.6857127934588849 * math.Pow(10,-16)
    actual = abraham_lorentz_force(up_quark_q, vacuum_permittivity,
      grav_jerk)

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
    var L float64 = 57909050.0
    var T float64 = 47.362
    var e float64 = 0.205630
    expected = 0.065884179454766778
    actual = perihelion_shift(L, T, e) * seconds_in_a_day * 59

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
    actual = schwarzschild_radius(mass_of_the_earth);

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
