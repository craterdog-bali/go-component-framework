/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"strings"
)

// NARRATIVE STRING INTERFACE

// This constructor attempts to create a new narrative string from the specified
// formatted string. It returns a narrative value and whether or not the string
// contained a valid narrative string.
// For valid string formats for this type see `../abstractions/language.go`.
func NarrativeFromString(v string) (Narrative, bool) {
	var ok = true
	var narrative string
	var matches = abstractions.ScanNarrative([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		narrative = trimTabs(matches[0])
	}
	return Narrative(narrative), ok
}

// This constructor attempts to create a new narrative string from the specified
// array of runes. It returns the corresponding narrative string.
func NarrativeFromRunes(v []rune) Narrative {
	var narrative, ok = NarrativeFromString(`"` + "\n" + string(v) + "\n" + `"`)
	if !ok {
		panic(fmt.Sprintf("The runes contain an illegal character: %v", string(v)))
	}
	return narrative
}

// This type defines the methods associated with a narrative string that extends
// the native Go string type and represents the string of runes that make up
// the narrative string.
type Narrative string

// LEXICAL INTERFACE

// This method returns the canonical string for this string.
func (v Narrative) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Narrative) String() string {
	return v.AsString()
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Narrative) IsEmpty() bool {
	return len(v) == 4 // The empty narrative string is: '"' EOL EOL '"'.
}

// This method returns the number of runes contained in this string.
func (v Narrative) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the runes in this string. The runes retrieved
// are in the same order as they are in the string.
func (v Narrative) AsArray() []rune {
	return []rune(v[2 : len(v)-2])
}

// INDEXED INTERFACE

// This method retrieves from this string the rune that is associated
// with the specified index.
func (v Narrative) GetItem(index int) rune {
	var runes = v.AsArray()
	var length = len(runes)
	index = abstractions.NormalizedIndex(index, length)
	return runes[index]
}

// This method retrieves from this string all runes from the first index
// through the last index (inclusive).
func (v Narrative) GetItems(first int, last int) []rune {
	var runes = v.AsArray()
	var length = len(runes)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	return runes[first : last+1]
}

// This method returns the index of the FIRST occurence of the specified rune
// in this string, or zero if this string does not contain the rune.
func (v Narrative) GetIndex(b rune) int {
	var runes = v.AsArray()
	for index, candidate := range runes {
		if candidate == b {
			// Found the rune.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The rune was not found.
	return 0
}

func trimTabs(v string) string {
	var result = `"` + "\n"
	var lines = strings.Split(v, "\n")
	var size = len(lines)
	var last = lines[size-1]  // The last line of the narrative should only be tabs.
	var tabs = len(last) - 1  // A count of the number of tabs in the last line.
	lines = lines[1 : size-1] // Trim off the first and last lines.
	for _, line := range lines {
		result += line[tabs:] + "\n" // Strip off the leading tabs.
	}
	result += `"`
	return result
}

// NARRATIVES LIBRARY

// This singleton creates a unique name space for the library functions for
// narratives.
var Narratives = &narratives{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for narratives.
type narratives struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified
// narrative strings.
func (l *narratives) Concatenate(first Narrative, second Narrative) Narrative {
	// TODO: This might need tweaking due to tabs before last double quote.
	var narrative = first[:len(first)-2] + second[2:]
	return Narrative(narrative)
}