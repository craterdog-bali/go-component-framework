/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	uti "github.com/bali-nebula/go-component-framework/utilities"
	col "github.com/craterdog/go-collection-framework/v2"
	reg "regexp"
)

// BYTECODE STRING IMPLEMENTATION

// These constants are used to define a regular expression for matching
// bytecode strings.
const (
	bytecode = `(` + base16 + `*)`
)

// This scanner is used for matching bytecode strings.
var bytecodeScanner = reg.MustCompile(`^(?:` + bytecode + `)$`)

// This constructor creates a new bytecode string from the specified string.
func BytecodeFromString(v string) abs.BytecodeLike {
	if !bytecodeScanner.MatchString(v) {
		var message = fmt.Sprintf("Attempted to construct a bytecode string from an invalid string: %v", v)
		panic(message)
	}
	return Bytecode(v)
}

// This constructor creates a new bytecode string from the specified instruction array.
// It returns the corresponding bytecode string.
func BytecodeFromArray(array []abs.Instruction) abs.BytecodeLike {
	var bytes = col.List[byte]()
	var instructions = col.ListFromArray(array)
	var iterator = col.Iterator[abs.Instruction](instructions)
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var firstByte = instruction.GetFirstByte()
		var secondByte = instruction.GetSecondByte()
		bytes.AddValue(firstByte)
		bytes.AddValue(secondByte)
	}
	var encoded = uti.Base16Encode(bytes.AsArray())
	return Bytecode(encoded)
}

// This constructor creates a new bytecode string from the specified byte
// sequence. It returns the corresponding bytecode string.
func BytecodeFromSequence(sequence abs.Sequential[abs.Instruction]) abs.BytecodeLike {
	var array = sequence.AsArray()
	return BytecodeFromArray(array)
}

// This type defines the methods associated with a bytecode string that extends
// the native Go string type and represents the string of instructions that make up
// the bytecode string.
type Bytecode string

// QUANTIZED INTERFACE

// This method returns a string value for this quantized string.
func (v Bytecode) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Bytecode) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of instructions contained in this string.
func (v Bytecode) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the instructions in this string. The instructions retrieved
// are in the same order as they are in the string.
func (v Bytecode) AsArray() []abs.Instruction {
	var encoded = string(v)
	var decoded = uti.Base16Decode(encoded)
	var bytes = col.ListFromArray(decoded)
	var instructions = col.List[abs.Instruction]()
	var iterator = col.Iterator[byte](bytes)
	for iterator.HasNext() {
		var firstByte = iterator.GetNext()
		var secondByte = iterator.GetNext()
		var instruction = abs.InstructionFromBytes(firstByte, secondByte)
		instructions.AddValue(instruction)
	}
	return instructions.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the instruction that is associated
// with the specified index.
func (v Bytecode) GetValue(index int) abs.Instruction {
	var array = v.AsArray()
	var instructions = col.Array[abs.Instruction](array)
	return instructions.GetValue(index)
}

// This method retrieves from this string all instructions from the first index
// through the last index (inclusive).
func (v Bytecode) GetValues(first int, last int) abs.Sequential[abs.Instruction] {
	var array = v.AsArray()
	var instructions = col.Array[abs.Instruction](array)
	return instructions.GetValues(first, last)
}
