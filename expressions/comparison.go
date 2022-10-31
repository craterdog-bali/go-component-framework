/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// COMPARISON EXPRESSION IMPLEMENTATION

// This constructor creates a new comparison expression.
func Comparison(first abs.ExpressionLike, operator abs.Operator, second abs.ExpressionLike) abs.ComparisonLike {
	var v = &comparisonExpression{}
	// Perform argument validation.
	v.SetFirst(first)
	v.SetOperator(operator)
	v.SetSecond(second)
	return v
}

// This type defines the structure and methods associated with an comparison
// expression.
type comparisonExpression struct {
	first    abs.ExpressionLike
	operator abs.Operator
	second   abs.ExpressionLike
}

// This method returns the first expression in this comparison expression.
func (v *comparisonExpression) GetFirst() abs.ExpressionLike {
	return v.first
}

// This method sets the first expression in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetFirst(first abs.ExpressionLike) {
	if first == nil {
		panic("The first expression in an comparison expression cannot be nil.")
	}
	v.first = first
}

// This method returns the comparison operator in this comparison expression.
func (v *comparisonExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the comparison operator in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetOperator(operator abs.Operator) {
	if operator < abs.LESS || operator > abs.MATCHES {
		panic("The operator in an comparison expression must be a valid comparison operation.")
	}
	v.operator = operator
}

// This method returns the second expression in this comparison expression.
func (v *comparisonExpression) GetSecond() abs.ExpressionLike {
	return v.second
}

// This method sets the second expression in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetSecond(second abs.ExpressionLike) {
	if second == nil {
		panic("The second expression in an comparison expression cannot be nil.")
	}
	v.second = second
}

// This method returns the type of this expression.
func (v *comparisonExpression) GetType() abs.Type {
	return abs.COMPARISON
}
