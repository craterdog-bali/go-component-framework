/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
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
	Key    any
	Value  any
	Entity any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all symbolic entities.
type Symbolic interface {
	GetIdentifier() string
}

// This interface defines the methods supported by all binding types.
// It binds a readonly key with a setable value.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This interface defines the methods supported by all sequences of values.
type Sequential[V Value] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

// This interface defines the methods supported by all sequences whose values can
// be indexed. The indices of an indexed sequence are ORDINAL rather than ZERO
// based (which is "SO last century"). This allows for positive indices starting
// at the beginning of the sequence, and negative indices starting at the end of
// the sequence as follows:
//
//	    1           2           3             N
//	[value 1] . [value 2] . [value 3] ... [value N]
//	   -N        -(N-1)      -(N-2)          -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Indexed[V Value] interface {
	ContainsValue(value V) bool
	GetValue(index int) V
	GetIndex(value V) int
}

// This interface defines the methods supported by all ratcheted agents that
// are capable of moving forward and backward over the values in a sequence. It
// is used to implement the GoF Iterator Pattern:
//   - https://en.wikipedia.org/wiki/Iterator_pattern
//
// A ratcheted agent locks into the slots that reside between each value in the
// sequence:
//
//	    [value 1] . [value 2] . [value 3] ... [value N]
//	  ^           ^           ^                         ^
//	slot 0      slot 1      slot 2                    slot N
//
// It moves from slot to slot and has access to the values (if they exist) on
// each side of the slot.
type Ratcheted[V Value] interface {
	GetSlot() int
	ToSlot(slot int)
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by iterator-like
// agents.
type IteratorLike[V Value] interface {
	Ratcheted[V]
}
