/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Element any
)

// INDIVIDUAL INTERFACES

type Complex interface {
	AsComplex() complex128
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() AngleLike
}

type Continuous interface {
	AsReal() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
}

type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

type Lexical interface {
	AsString() string
}

type Matchable interface {
	MatchesText(text string) bool
	GetMatches(text string) []string
}

type Polarized interface {
	IsNegative() bool
}

type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

type Temporal interface {
	// Return the entire time in specific units.
	AsMilliseconds() float64
	AsSeconds() float64
	AsMinutes() float64
	AsHours() float64
	AsDays() float64
	AsWeeks() float64
	AsMonths() float64
	AsYears() float64

	// Return a specific part of the entire time.
	GetMilliseconds() int
	GetSeconds() int
	GetMinutes() int
	GetHours() int
	GetDays() int
	GetWeeks() int
	GetMonths() int
	GetYears() int
}

// CONSOLIDATED INTERFACES

type AngleLike interface {
	Continuous
}

type BooleanLike interface {
	Discrete
}

type CharacterLike interface {
	Discrete
}

type DurationLike interface {
	Discrete
	Polarized
	Temporal
}

type IntegerLike interface {
	Discrete
	Polarized
}

type MomentLike interface {
	Discrete
	Temporal
}

type NumberLike interface {
	Continuous
	Polarized
	Complex
}

type PatternLike interface {
	Lexical
	Matchable
}

type PercentageLike interface {
	Discrete
	Continuous
	Polarized
}

type ProbabilityLike interface {
	Discrete
	Continuous
}

type RealLike interface {
	Continuous
	Polarized
}

type ResourceLike interface {
	Lexical
	Segmented
}
