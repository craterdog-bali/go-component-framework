/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"reflect"
)

// COMPONENT IMPLEMENTATION

// This constructor creates a new component.
func Component[T any](entity T) abstractions.ComponentLike[T] {
	var v = &component[T]{}
	// Perform argument validation.
	v.SetEntity(entity)
	return v
}

// This constructor creates a new component with the specified context.
func ComponentWithContext[T any](entity T, context abstractions.ContextLike) abstractions.ComponentLike[T] {
	var v = &component[T]{}
	// Perform argument validation.
	v.SetEntity(entity)
	v.SetContext(context)
	return v
}

// This type defines the structure and methods associated with a generic
// component.
// This type is parameterized as follows:
//   - T is any type of entity.
type component[T any] struct {
	entity  T
	context abstractions.ContextLike
	note    string
}

// COMPONENT IMPLEMENTATION

// This method determines whether or not this component is parameterized.
func (v *component[T]) IsGeneric() bool {
	return v.context != nil
}

func (v *component[T]) IsAnnotated() bool {
	return len(v.note) > 0
}

// This method returns the entity for this component.
func (v *component[T]) GetEntity() T {
	return v.entity
}

// This method sets the entity for this component.
func (v *component[T]) SetEntity(entity T) {
	if !reflect.ValueOf(entity).IsValid() {
		panic("A component requires an entity.")
	}
	v.entity = entity
}

// This method returns the context for this component.
func (v *component[T]) GetContext() abstractions.ContextLike {
	return v.context
}

// This method sets the context for this component.
func (v *component[T]) SetContext(context abstractions.ContextLike) {
	v.context = context
}

// This method returns the note for this component.
func (v *component[T]) GetNote() string {
	return v.note
}

// This method sets the note for this component.
func (v *component[T]) SetNote(note string) {
	v.note = note
}