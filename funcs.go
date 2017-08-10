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
    "math"
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

