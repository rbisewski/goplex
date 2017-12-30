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
func tsiolkovskyDeltaV(Ve float64, m0 float64, mf float64) float64 {

	// input validation
	if mf == 0 {
		return 0
	}

	// calculate the mass ratio, i.e. the different between the initial
	// propellant and the final mass w/o propellant
	ratioOfInitialToDryMass := m0 / mf

	// determine the natural log of the mass ratio
	nlogOfMassRatio := math.Log(ratioOfInitialToDryMass)

	// figure out the total energy requiring during the change of mass
	// over the start and end of the launch, otherwise known as the delta-V
	deltaV := Ve * nlogOfMassRatio

	// go ahead and return the values
	return deltaV
}

//! Function to calculate the energy of a photon
/*
 * @param    float64    wavelength --> l
 *
 * @result   float64    energy of a photon, in Joules
 */
func photonEnergy(l float64) float64 {

	// if wavelength is zero, return zero
	if l == 0 {
		return 0
	}

	// compare the wavelength to the planck constant w/ speed of light
	// and then obtain the ratio of that to the wavelength
	return planckConstant * c / l
}

//! Thermal velocity of a heated gas
/*
 * @param    float64    gravity acceleration at sea-level --> g
 * @param    float64    temperature, in Kelvins           --> T
 * @param    float64    mass of exhaust, per molecule     --> m
 *
 * @result   float64    velocity of the gas in question
 */
func thermalVelocityOfHeatedGas(g float64, T float64,
	m float64) float64 {

	// input validation
	if T < 0 || g == 0 || m == 0 {
		return 0
	}

	// inverse of the acceleration
	inverseG := 1 / g

	// ratio of the boltzmann & temperature to the molecular mass
	boltzRatioToMass := 3 * boltzmannConstantEv * T / m

	// since this deals with fluid dynamics in space, take the square of
	// the boltz-mass ratio
	squareOfBoltzRatio := math.Sqrt(boltzRatioToMass)

	// apply the inverse g, which yields the specific impulse
	specificImpulse := inverseG * squareOfBoltzRatio

	// go ahead and return the thermal velocity
	return specificImpulse
}

//! Function to calculate the relativistic doppler effect
/*
 * @param    float64    velocity   --> v
 *
 * @result   float64    time dilation ratio
 */
func lorentzFactor(v float64) float64 {

	// ensure that the velocity is not equal to c
	if v == c {
		return 0.0
	}

	// determine the square factor
	squareFactor := 1 - ((v * v) / (c * c))

	// take the square root of the factor
	sqrtFactor := math.Sqrt(squareFactor)

	// safety check, ensure that that factor is not zero
	if sqrtFactor == 0.0 {
		return 0.0
	}

	// go ahead and return the inverse square root
	return 1 / sqrtFactor
}

//! Function to calculate the Abraham-Lorentz force
/*
 * @param    float64    charge            --> q
 * @param    float64    electric constant --> e0
 * @param    float64    jerk              --> a
 *
 * @result   float64    time dilation ratio
 */
func abrahamLorentzForce(q float64, e0 float64, a float64) float64 {

	// ensure that the electrical constant isn't zero
	if e0 == 0.0 {
		return 0.0
	}

	// calculate the square of the charge
	squareOfCharge := q * q

	// calculate the charged field value, for a vacuum
	chargedFieldValue := 6 * math.Pi * e0 * c * c * c

	// calculate the ratio of the charge to the value of the field
	ratioOfChargeToField := squareOfCharge / chargedFieldValue

	// adjust the field with the given jerk value
	return ratioOfChargeToField * a
}

//! Function to calculate the perihelion shift of an orbit
/*
 * @param    float64    semi-major axis      --> L
 * @param    float64    orbital speed        --> T
 * @param    float64    orbital eccentricity --> e
 *
 * @result   float64    perihelion shift, in radians/revolution
 */
func perihelionShift(L float64, T float64, e float64) float64 {

	// speed of light in kilometers per second
	cInKmPerSecond := c * 1000.0

	// calculate the spherical shape of the semi-major axis
	dividend := 24 * math.Pi * math.Pi * math.Pi * L * L

	// calculate the velocity of the orbital body
	divisor := T * T * cInKmPerSecond * cInKmPerSecond * (1 - e*e)

	// safety check, if the divisor is zero, return 0
	if divisor == 0.0 {
		return 0.0
	}

	// combine both to obtain the perihelion shift
	return dividend / divisor
}

//! Function to calculate the Schwarzschild radius for a given mass,
//! i.e. sphere escape velocity at the speed of light.
/*
 * @param    float64    mass --> M
 *
 * @result   float64    Schwarzschild radius, in units
 */
func schwarzschildRadius(M float64) float64 {

	// input validation
	if M <= 0.0 {
		return 0.0
	}

	// speed of light in a vacuum, squared
	speedOfLightInVacSquared := c * c

	// ratio of mass to gravity, as per the universal constant
	ratioOfMassToGravity := 2 * universalGravitationConstant * M

	// pass back the calcuated value
	return ratioOfMassToGravity / speedOfLightInVacSquared
}
