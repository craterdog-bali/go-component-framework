/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	col "github.com/craterdog/go-collection-framework"
)

// STACK IMPLEMENTATION

// This constructor creates a new empty stack with the default capacity.
// The default capacity is 16 values.
func Stack() abs.StackLike {
	return col.Stack[abs.ComponentLike]()
}

// This constructor creates a new empty stack with the specified capacity.
func StackWithCapacity(capacity int) abs.StackLike {
	return col.StackWithCapacity[abs.ComponentLike](capacity)
}
