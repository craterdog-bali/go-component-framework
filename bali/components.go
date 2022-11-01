/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	com "github.com/craterdog-bali/go-bali-document-notation/components"
	sts "strings"
)

// PARSING METHODS
//
// Each parsing method returns the following three results:
//  * The parsed component corresponding to the source string, or nil if the
//    parsing failed.
//  * The parsed token that caused the failure if the parsing did fail.
//  * Whether or not the parsing succeeded.

// This method attempts to parse annotation. It returns the annotation and
// whether or not the annotation was successfully parsed.
func (v *parser) parseAnnotation() (string, *Token, bool) {
	var ok bool
	var token *Token
	var annotation string
	annotation, token, ok = v.parseNote()
	if !ok {
		annotation, token, ok = v.parseComment()
	}
	return annotation, token, ok
}

// This method attempts to parse a comment. It returns a string containing the
// comment and whether or not the comment was successfully parsed.
func (v *parser) parseComment() (string, *Token, bool) {
	var comment string
	var token = v.nextToken()
	if token.Type != TokenComment {
		v.backupOne()
		return comment, token, false
	}
	comment = trimIndentation(token.Value)
	return comment, token, true
}

// This method adds the canonical format for the specified annotation to the
// state of the formatter.
func (v *formatter) formatComment(comment string) {
	var lines = sts.Split(comment, "\n")
	v.state.AppendString(`!>`)
	for _, line := range lines {
		v.state.AppendNewline()
		v.state.AppendString(line)
	}
	v.state.AppendNewline()
	v.state.AppendString(`<!`)
}

// This method attempts to parse a component. It returns the component and
// whether or not the component was successfully parsed.
func (v *parser) parseComponent() (abs.ComponentLike, *Token, bool) {
	var component abs.ComponentLike
	var context abs.ContextLike
	var note string
	var entity, token, ok = v.parseEntity()
	if ok {
		context, token, _ = v.parseContext() // The context is optional.
		note, token, _ = v.parseNote()       // The note is optional.
		component = com.ComponentWithContext(entity, context)
		component.SetNote(note)
	}
	return component, token, ok
}

// This method adds the canonical format for the specified component to the
// state of the formatter.
func (v *formatter) formatComponent(component abs.ComponentLike) {
	var entity = component.GetEntity()
	v.formatAny(entity)
	if component.IsGeneric() {
		var context = component.GetContext()
		v.formatContext(context)
	}
	if component.IsAnnotated() {
		v.state.AppendString("  ")
		var note = component.GetNote()
		v.formatNote(note)
	}
}

// This method attempts to parse the context for a parameterized component. It
// returns an array of parameters and whether or not the context was
// successfully parsed.
func (v *parser) parseContext() (abs.ContextLike, *Token, bool) {
	var ok bool
	var token *Token
	var parameters abs.CatalogLike[abs.Symbolic, abs.ComponentLike]
	var context abs.ContextLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		return nil, token, false
	}
	parameters, token, ok = v.parseParameters()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$parameter",
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar(")",
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	context = com.Context(parameters)
	return context, token, true
}

// This method adds the canonical format for the specified context to the
// state of the formatter.
func (v *formatter) formatContext(context abs.ContextLike) {
	v.state.AppendString("(")
	var parameters = context.GetParameters()
	v.formatParameters(parameters)
	v.state.AppendString(")")
}

// This method attempts to parse the specified delimiter. It returns
// the token and whether or not the delimiter was found.
func (v *parser) parseDelimiter(delimiter string) (string, *Token, bool) {
	var token = v.nextToken()
	if token.Type == TokenEOF || token.Value != delimiter {
		v.backupOne()
		return delimiter, token, false
	}
	return delimiter, token, true
}

// This method attempts to parse a component entity. It returns the component
// entity and whether or not the component entity was successfully parsed.
func (v *parser) parseEntity() (abs.EntityLike, *Token, bool) {
	var ok bool
	var token *Token
	var entity abs.EntityLike
	entity, token, ok = v.parseElement()
	if !ok {
		entity, token, ok = v.parseString()
	}
	if !ok {
		entity, token, ok = v.parseCollection()
	}
	if !ok {
		entity, token, ok = v.parseProcedure()
	}
	return entity, token, ok
}

// This method attempts to parse the end-of-file (EOF) marker. It returns
// the token and whether or not an EOL token was found.
func (v *parser) parseEOF() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOF {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse the end-of-line (EOL) marker. It returns
// the token and whether or not an EOF token was found.
func (v *parser) parseEOL() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOL {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseIdentifier() (string, *Token, bool) {
	var identifier string = "<UNKNOWN>"
	var token = v.nextToken()
	if token.Type != TokenIdentifier {
		v.backupOne()
		return identifier, token, false
	}
	identifier = token.Value
	return identifier, token, true
}

// This method adds the canonical format for the specified identifier to the
// state of the formatter.
func (v *formatter) formatIdentifier(identifier string) {
	v.state.AppendString(identifier)
}

// This method attempts to parse the specified keyword. It returns
// the token and whether or not the keyword was found.
func (v *parser) parseKeyword(keyword string) (string, *Token, bool) {
	var token = v.nextToken()
	if token.Type == TokenKeyword || token.Value != keyword {
		v.backupOne()
		return keyword, token, false
	}
	return keyword, token, true
}

// This method attempts to parse a note. It returns a string containing the
// note and whether or not the note was successfully parsed.
func (v *parser) parseNote() (string, *Token, bool) {
	var note string
	var token = v.nextToken()
	if token.Type != TokenNote {
		v.backupOne()
		return note, token, false
	}
	note = token.Value[2:] // Remove the leading "! ".
	return note, token, true
}

// This method adds the canonical format for the specified annotation to the
// state of the formatter.
func (v *formatter) formatNote(note string) {
	v.state.AppendString("! ")
	v.state.AppendString(note)
}

// This method attempts to parse a parameter containing a name and value. It
// returns the parameter and whether or not the parameter was successfully
// parsed.
func (v *parser) parseParameter() (abs.AssociationLike[abs.Symbolic, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var name abs.Symbolic
	var value abs.ComponentLike
	name, token, ok = v.parseSymbol()
	if !ok {
		// This is not a parameter.
		return nil, token, false
	}
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		// This is not a parameter.
		v.backupOne() // Put back the symbol token.
		return nil, token, false
	}
	// This must be a parameter.
	value, token, ok = v.parseComponent()
	if !ok {
		panic("Expected a value after the ':' character.")
	}
	var parameter = col.Association[abs.Symbolic, abs.ComponentLike](name, value)
	return parameter, token, true
}

// This method adds the canonical format for the specified parameter to the
// state of the formatter.
func (v *formatter) formatParameter(parameter abs.AssociationLike[abs.Symbolic, abs.ComponentLike]) {
	var key = parameter.GetKey()
	v.state.AppendString("$")
	v.formatIdentifier(key.GetIdentifier())
	v.state.AppendString(": ")
	var value = parameter.GetValue()
	v.formatComponent(value)
}

// This method attempts to parse context parameters. It returns the
// context parameters and whether or not the context parameters were
// successfully parsed.
func (v *parser) parseParameters() (abs.CatalogLike[abs.Symbolic, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var parameter abs.AssociationLike[abs.Symbolic, abs.ComponentLike]
	var parameters = col.Catalog[abs.Symbolic, abs.ComponentLike]()
	_, token, ok = v.parseEOL()
	if !ok {
		// The parameters are on a single line.
		_, token, ok = v.parseDelimiter(":")
		if ok {
			panic("There must be at least one parameter in a context.")
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			panic("There must be at least one parameter in a context.")
		}
		for {
			parameters.AddAssociation(parameter)
			// Every subsequent parameter must be preceded by a ','.
			_, token, ok = v.parseDelimiter(",")
			if !ok {
				// No more parameters.
				break
			}
			parameter, token, ok = v.parseParameter()
			if !ok {
				panic("Expected a parameter after the ',' character.")
			}
		}
		return parameters, token, true
	}
	// The parameters are on separate lines.
	parameter, token, ok = v.parseParameter()
	if !ok {
		panic("There must be at least one parameter in a context.")
	}
	for {
		parameters.AddAssociation(parameter)
		// Every parameter must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the parameter.")
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			// No more parameters.
			break
		}
	}
	return parameters, token, true
}

// This method adds the canonical format for the specified parameters to the
// state of the formatter.
func (v *formatter) formatParameters(parameters abs.CatalogLike[abs.Symbolic, abs.ComponentLike]) {
	switch parameters.GetSize() {
	case 0:
		panic("A context must have at least one parameter.")
	case 1:
		var parameter = parameters.GetItem(1)
		v.formatParameter(parameter)
	default:
		var iterator = age.Iterator[abs.AssociationLike[abs.Symbolic, abs.ComponentLike]](parameters)
		v.state.IncrementDepth()
		for iterator.HasNext() {
			v.state.AppendNewline()
			var parameter = iterator.GetNext()
			v.formatParameter(parameter)
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	}
}

// PRIVATE FUNCTIONS

// This function removes the first and last line delimiters (shown as "xx")
// and the indentation from each line of the specified multi-line string.
//
// The following is an indented string with dashes showing the indentation:
//
//	----xx
//	----This is the first line
//	----of a multi-line
//	----indented string.
//	----xx
//
// It will be trimmed to:
//
//	This is the first line
//	of a multi-line
//	indented string.
//
// Only the indented lines remain and are no longer indented.
func trimIndentation(v string) string {
	var trimmed string
	var lines = sts.Split(v, "\n")
	var size = len(lines) - 1
	var last = lines[size]          // The last line provides the level of indentation.
	var indentation = len(last) - 2 // The number of spaces in the last line.
	lines = lines[1:size]           // Strip off the first and last delimitier lines.
	for _, line := range lines {
		trimmed += line[indentation:] + "\n" // Strip off the indentation.
	}
	return trimmed[:len(trimmed)-1] // Strip off the extra end-of-line character.
}
