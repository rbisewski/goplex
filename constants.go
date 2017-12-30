/*
 * Goplex Math Constants
 *
 * Description: A set of common mathematical and physics related constants
 *              that overlap with many scientific fields of study.
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

//
// Globals
//
var (

	// seconds in a single Earth day
	secondsInADay = 86400.0

	// Speed of light in a vaccuum, in m/s
	c = 299792458.0

	// Universal gravitational constant, in m^3 kg^-1 s^-2
	universalGravitationConstant = 0.0000000000667408

	// Planck constant, in Joule seconds
	planckConstant = 6.626069934 * math.Pow(10, -34)

	// Boltzmann constant, in Joules per Kelvin
	boltzmannConstantJoules = 1.38064852 * math.Pow(10, -23)

	// Boltzmann constant, in eV per Kelvin
	boltzmannConstantEv = 8.6173303 * math.Pow(10, -5)

	// Vacuum permittivity, in Farads per metre
	vacuumPermittivity = 8.854187817 * math.Pow(10, -12)

	// Mass of the planet Earth, in kilograms
	massOfTheEarth = 5.97237 * math.Pow(10, 24)
)
