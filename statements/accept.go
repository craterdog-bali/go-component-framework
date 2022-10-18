/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package statements

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ACCEPT CLAUSE IMPLEMENTATION

// This constructor creates a new accept clause.
func Accept(message any) abstractions.AcceptLike {
	var v = &acceptClause{}
	// Perform argument validation.
	v.SetMessage(message)
	return v
}

// This type defines the structure and methods associated with an accept
// clause.
type acceptClause struct {
	message any
}

// This method returns the message expression for this accept clause.
func (v *acceptClause) GetMessage() any {
	return v.message
}

// This method sets the message expression for this accept clause.
func (v *acceptClause) SetMessage(message any) {
	if message == nil {
		panic("An accept clause requires a message.")
	}
	v.message = message
}