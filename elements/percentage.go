/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"strconv"
)

// PERCENTAGE INTERFACE

// This constructor attempts to create a new percentage from the specified
// formatted string. It returns a percentage value and whether or not the string
// contained a valid percentage.
// For valid string formats for this type see `../abstractions/language.go`.
func PercentageFromString(v string) (Percentage, bool) {
	var percentage, ok = stringToPercentage(v)
	return Percentage(percentage), ok
}

// This type defines the methods associated with percentage elements. It extends
// the native Go float64 type and represents a percentage. Percentages can be
// negative.
type Percentage float64

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Percentage) AsString() string {
	return strconv.FormatFloat(float64(v), 'G', -1, 64) + "%"
}

// This method implements the standard Go Stringer interface.
func (v Percentage) String() string {
	return v.AsString()
}

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous component.
func (v Percentage) AsReal() float64 {
	return float64(v / 100.0)
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Percentage) IsNegative() bool {
	return v < 0.0
}

// PERCENTAGES LIBRARY

// This singleton creates a unique name space for the library functions for
// percentage elements.
var Percentages = &percentages{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for percentage elements.
type percentages struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified percentage.
func (l *percentages) Inverse(percentage Percentage) Percentage {
	return Percentage(float64(-percentage))
}

// This library function returns the sum of the specified percentages.
func (l *percentages) Sum(first, second Percentage) Percentage {
	return Percentage(float64(first + second))
}

// This library function returns the difference of the specified percentages.
func (l *percentages) Difference(first, second Percentage) Percentage {
	return Percentage(float64(first - second))
}

// This library function returns the specified percentage scaled by the
// specified factor.
func (l *percentages) Scaled(percentage Percentage, factor float64) Percentage {
	return Percentage(float64(percentage) * factor)
}

// PRIVATE FUNCTIONS

// This function parses a percentage string and returns the corresponding
// floating point number and whether or not the string contained a valid
// percentage.
func stringToPercentage(v string) (float64, bool) {
	var percentage float64
	var ok = true
	var matches = abstractions.ScanPercentage([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		var err error
		percentage, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
			ok = false
		}
	}
	return percentage, ok
}