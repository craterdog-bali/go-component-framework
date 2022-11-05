/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ACCEPT CLAUSE IMPLEMENTATION

// This constructor creates a new accept clause.
func AcceptClause(message abs.ExpressionLike) abs.AcceptClauseLike {
	var v = &acceptClause{}
	// Perform argument validation.
	v.SetMessage(message)
	return v
}

// This type defines the structure and methods associated with an accept
// clause.
type acceptClause struct {
	message abs.ExpressionLike
}

// This method is a dummy method that always returns true.
func (v *acceptClause) IsAcceptClause() bool {
	return true
}

// This method returns the message expression for this accept clause.
func (v *acceptClause) GetMessage() abs.ExpressionLike {
	return v.message
}

// This method sets the message expression for this accept clause.
func (v *acceptClause) SetMessage(message abs.ExpressionLike) {
	if message == nil {
		panic("An accept clause requires a message.")
	}
	v.message = message
}
