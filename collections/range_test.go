/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRangesWithIntegers(t *tes.T) {
	var s = col.Range[int](3, "..", 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.Equal(t, 5, s.GetItem(3))
	ass.Equal(t, 0, s.GetIndex(2))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 3, s.GetIndex(5))
	ass.Equal(t, 5, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []int{3, 4, 5, 6, 7}, s.AsArray())
	var iterator = age.Iterator[int](s)
	ass.Equal(t, 3, iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, 7, iterator.GetPrevious())

	s = col.Range[int](3, "<..", 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.Equal(t, 4, s.GetItem(1))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 2, s.GetIndex(5))
	ass.Equal(t, 4, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []int{4, 5, 6, 7}, s.AsArray())

	s = col.Range[int](3, "<..<", 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 3, s.GetSize())
	ass.Equal(t, 4, s.GetItem(1))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 3, s.GetIndex(6))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, []int{4, 5, 6}, s.AsArray())

	s = col.Range[int](3, "..<", 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.Equal(t, 3, s.GetItem(1))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 4, s.GetIndex(6))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, []int{3, 4, 5, 6}, s.AsArray())
}

func TestRangesWithRunes(t *tes.T) {
	var s = col.Range[rune]('a', "..", 'z')
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 26, s.GetSize())
	ass.Equal(t, 'd', s.GetItem(4))
	ass.Equal(t, 13, s.GetIndex('m'))

	var a = col.List[rune]()
	var items = s.AsArray()
	a.AddItems(s)
	ass.Equal(t, items, a.AsArray())
}
