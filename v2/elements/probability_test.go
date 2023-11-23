/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBooleanProbabilities(t *tes.T) {
	var v1 = ele.Probability.FromBool(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = ele.Probability.FromBool(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = ele.Probability.FromFloat(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = ele.Probability.FromFloat(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbabilities(t *tes.T) {
	ele.Probability.Random()
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = ele.Probability.FromFloat(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = ele.Probability.FromFloat(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = ele.Probability.FromFloat(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = ele.Probability.FromFloat(0.75)
	var F = ele.Probability.FromFloat(0.25)

	var andNot = ele.Probability.And(ele.Probability.Not(T), ele.Probability.Not(T))
	var notOr = ele.Probability.Not(ele.Probability.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probability.And(ele.Probability.Not(T), ele.Probability.Not(F))
	notOr = ele.Probability.Not(ele.Probability.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probability.And(ele.Probability.Not(F), ele.Probability.Not(T))
	notOr = ele.Probability.Not(ele.Probability.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probability.And(ele.Probability.Not(F), ele.Probability.Not(F))
	notOr = ele.Probability.Not(ele.Probability.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = ele.Probability.And(T, ele.Probability.Not(T))
	ass.Equal(t, sans, ele.Probability.Sans(T, T))

	sans = ele.Probability.And(T, ele.Probability.Not(F))
	ass.Equal(t, sans, ele.Probability.Sans(T, F))

	sans = ele.Probability.And(F, ele.Probability.Not(T))
	ass.Equal(t, sans, ele.Probability.Sans(F, T))

	sans = ele.Probability.And(F, ele.Probability.Not(F))
	ass.Equal(t, sans, ele.Probability.Sans(F, F))

	var xor = ele.Probability.FromFloat(ele.Probability.Sans(T, T).AsFloat() + ele.Probability.Sans(T, T).AsFloat())
	ass.Equal(t, xor, ele.Probability.Xor(T, T))

	xor = ele.Probability.FromFloat(ele.Probability.Sans(T, F).AsFloat() + ele.Probability.Sans(F, T).AsFloat())
	ass.Equal(t, xor, ele.Probability.Xor(T, F))

	xor = ele.Probability.FromFloat(ele.Probability.Sans(F, T).AsFloat() + ele.Probability.Sans(T, F).AsFloat())
	ass.Equal(t, xor, ele.Probability.Xor(F, T))

	xor = ele.Probability.FromFloat(ele.Probability.Sans(F, F).AsFloat() + ele.Probability.Sans(F, F).AsFloat())
	ass.Equal(t, xor, ele.Probability.Xor(F, F))
}
