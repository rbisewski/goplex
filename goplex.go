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
    // Energy of a photon with a wavelength of 400nm
    //
    var photon_energy_expected_result float64 = 4.966114480984394 *
      math.Pow(10, -19)
    var wavelength float64 = 400 * math.Pow(10, -9)
    var photon_energy_test = photon_energy(wavelength)

    // test to ensure this got the expected result
    if photon_energy_expected_result != photon_energy_test {
        fmt.Println("Photon Energy test failed!")
        fmt.Println("Expected: ", photon_energy_expected_result)
        fmt.Println("Calculated: ", photon_energy_test)
        os.Exit(1)
    }

    //
    // Thermal velocity of gas propellant
    //
    var thermal_velocity_of_gas_expected_result float64 = 0.11043619735553113
    var g float64               = 9.8000
    var m float64               = 0.8100     // RP-1 density
    var temp_in_kelvins float64 = 3670.0000  // RP-1 temp
    var thermal_velocity_of_gas_test float64 = thermal_velocity_of_heated_gas(g,
      temp_in_kelvins, m)

    // test to ensure this got the expected result
    if thermal_velocity_of_gas_expected_result != thermal_velocity_of_gas_test {
        fmt.Println("Thermal velocity test failed!")
        fmt.Println("Expected: ", thermal_velocity_of_gas_expected_result)
        fmt.Println("Calculated: ", thermal_velocity_of_gas_test)
        os.Exit(1)
    }

    //
    // Calculate the Lorentz factor of 0.5c
    //
    var lorentz_factor_expected float64 = 1.1547005383792517
    var half_c float64                  = c / 2
    var lorentz_factor_test float64     = lorentz_factor(half_c)

    // test to ensure this got the expected result
    if lorentz_factor_expected != lorentz_factor_test {
        fmt.Println("Lorentz factor test failed!")
        fmt.Println("Expected: ", lorentz_factor_expected)
        fmt.Println("Calculated: ", lorentz_factor_test)
        os.Exit(1)
    }

    //
    // Abraham-Lorentz force for a up quark in a vacuum
    //
    var up_force_present_in_vacuum_expected float64 = 9.6857127934588849 *
      math.Pow(10,-16)
    var up_quark_q float64 = 0.66666666
    var grav_jerk float64 = 9.8
    var up_force_present_in_vacuum_test float64 =
      abraham_lorentz_force(up_quark_q, vacuum_permittivity, grav_jerk)

    // test to ensure this got the expected result
    if up_force_present_in_vacuum_expected != up_force_present_in_vacuum_test {
        fmt.Println("Abraham-Lorentz test failed!")
        fmt.Println("Expected: ", up_force_present_in_vacuum_expected)
        fmt.Println("Calculated: ", up_force_present_in_vacuum_test)
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
