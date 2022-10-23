/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const bad = `">
<"`

const n0 = `">

<"`

const n1 = `">
	abcd本1234
<"`

const n2 = `">
d本1
<"`

const n3 = `">
	abcd本
<"`

const n4 = `">
1234
<"`

func TestEmptyNarrative(t *tes.T) {
	var v = str.Narrative(n0)
	ass.Equal(t, n0, lan.FormatValue(v))
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestNarrative(t *tes.T) {
	var v = str.Narrative(n1)
	ass.Equal(t, n1, lan.FormatValue(v))
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 10, v.GetSize())
	ass.Equal(t, 'a', v.GetItem(2))
	ass.Equal(t, '4', v.GetItem(-1))
	ass.Equal(t, n2, lan.FormatValue(str.Narrative(string(v.GetItems(5, 7)))))
	ass.Equal(t, 6, v.GetIndex('本'))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = str.Narrative(n3)
	var v2 = str.Narrative(n4)
	ass.Equal(t, n1, lan.FormatValue(str.Narratives.Concatenate(v1, v2)))
}
