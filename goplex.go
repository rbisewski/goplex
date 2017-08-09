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

    // Planck constant, in Joule seconds
    planck_constant = 6.626069934 * math.Pow(10, -34)

    // Boltzmann constant, in Joules per Kelvin
    boltzmann_constant_joules = 1.38064852 * math.Pow(10, -23)

    // Boltzmann constant, in eV per Kelvin
    boltzmann_constant_ev = 8.6173303 * math.Pow(10, -5)

    // Vacuum permittivity, in Farads per metre
    vacuum_permittivity = 8.854187817 * math.Pow(10, -12)
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

//! Function to calculate the energy of a photon
/*
 * @param    float64    wavelength --> l
 *
 * @result   float64    energy of a photon, in Joules
 */
func photon_energy(l float64) (float64) {

    // if wavelength is zero, return zero
    if l == 0 {
        return 0
    }

    // compare the wavelength to the planck constant w/ speed of light
    // and then obtain the ratio of that to the wavelength
    return planck_constant * c / l
}

//! Thermal velocity of a heated gas
/*
 * @param    float64    gravity acceleration at sea-level --> g
 * @param    float64    temperature, in Kelvins           --> T
 * @param    float64    mass of exhaust, per molecule     --> m
 *
 * @result   float64    velocity of the gas in question
 */
func thermal_velocity_of_heated_gas(g float64, T float64,
  m float64) (float64) {

    // input validation
    if T < 0 || g == 0 || m == 0 {
        return 0
    }

    // inverse of the acceleration
    var inverse_g float64 = 1 / g

    // ratio of the boltzmann & temperature to the molecular mass
    var boltz_ratio_to_mass float64 = 3 * boltzmann_constant_ev * T / m

    // since this deals with fluid dynamics in space, take the square of
    // the boltz-mass ratio
    var square_of_boltz_ratio float64 = math.Sqrt(boltz_ratio_to_mass)

    // apply the inverse g, which yields the specific impulse
    var specific_impulse float64 = inverse_g * square_of_boltz_ratio

    // go ahead and return the thermal velocity
    return specific_impulse
}

//! Function to calculate the relativistic doppler effect
/*
 * @param    float64    velocity   --> v
 *
 * @result   float64    time dilation ratio
 */
func lorentz_factor(v float64) (float64) {

    // ensure that the velocity is not equal to c
    if v == c {
        return 0.0
    }

    // determine the square factor
    var square_factor float64 = 1 - ((v * v) / (c * c))

    // take the square root of the factor
    var sqrt_factor float64 = math.Sqrt(square_factor)

    // safety check, ensure that that factor is not zero
    if sqrt_factor == 0.0 {
        return 0.0
    }

    // go ahead and return the inverse square root
    return 1 / sqrt_factor
}

//! Function to calculate the Abraham-Lorentz force
/*
 * @param    float64    charge            --> q
 * @param    float64    electric constant --> e0
 * @param    float64    jerk              --> a
 *
 * @result   float64    time dilation ratio
 */
func abraham_lorentz_force(q float64, e0 float64, a float64) (float64) {

    // ensure that the electrical constant isn't zero
    if e0 == 0.0 {
        return 0.0
    }

    // calculate the square of the charge
    var square_of_charge float64 = q * q

    // calculate the charged field value, for a vacuum
    var charged_field_value float64 = 6 * math.Pi * e0 * c * c * c

    // calculate the ratio of the charge to the value of the field
    var ratio_of_charge_to_field = square_of_charge / charged_field_value

    // adjust the field with the given jerk value
    return ratio_of_charge_to_field * a
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
