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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
)

// THROW CLAUSE IMPLEMENTATION

// This constructor creates a new throw clause.
func ThrowClause(exception abs.Expression) abs.ThrowClauseLike {
	var v = &throwClause{}
	// Perform argument validation.
	v.SetException(exception)
	return v
}

// This type defines the structure and methods associated with an throw
// clause.
type throwClause struct {
	exception abs.Expression
}

// This method is a dummy method that always returns true.
func (v *throwClause) IsThrowClause() bool {
	return true
}

// This method returns the exception expression for this throw clause.
func (v *throwClause) GetException() abs.Expression {
	return v.exception
}

// This method sets the exception expression for this throw clause.
func (v *throwClause) SetException(exception abs.Expression) {
	if exception == nil {
		panic("A throw clause requires an exception.")
	}
	v.exception = exception
}
