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
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	pro "github.com/craterdog-bali/go-bali-document-notation/procedures"
)

// This method attempts to parse an accept clause. It returns the accept
// clause and whether or not the accept clause was successfully parsed.
func (v *parser) parseAcceptClause() (abs.AcceptClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message abs.ExpressionLike
	var clause abs.AcceptClauseLike
	_, token, ok = v.parseKeyword("accept")
	if !ok {
		// This is not a accept clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$message",
			"$acceptClause",
			"$message")
		panic(message)
	}
	clause = pro.AcceptClause(message)
	return clause, token, true
}

// This method adds the canonical format for the specified accept clause to the
// state of the formatter.
func (v *formatter) formatAcceptClause(clause abs.AcceptClauseLike) {
	v.state.AppendString("accept ")
	var message = clause.GetMessage()
	v.formatExpression(message)
}

// This method attempts to parse an attribute. It returns the attribute and
// whether or not the attribute was successfully parsed.
func (v *parser) parseAttribute() (abs.AttributeLike, *Token, bool) {
	var ok bool
	var token *Token
	var variable string
	var indices = col.List[abs.ExpressionLike]()
	var attribute abs.AttributeLike
	variable, token, ok = v.parseIdentifier()
	if !ok {
		// This is not an attribute.
		return attribute, token, false
	}
	indices, token, ok = v.parseIndices()
	if !ok {
		// This is not an attribute.
		v.backupOne() // Put back the identifier token.
		return attribute, token, false
	}
	attribute = pro.Attribute(variable, indices)
	return attribute, token, true
}

// This method adds the canonical format for the specified attribute to the
// state of the formatter.
func (v *formatter) formatAttribute(attribute abs.AttributeLike) {
	var variable = attribute.GetVariable()
	v.state.AppendString(variable)
	var indices = attribute.GetIndices()
	v.formatIndices(indices)
}

// This method attempts to parse a do block. It returns the do block and whether
// or not the do block was successfully parsed.
func (v *parser) parseBlock() (abs.BlockLike, *Token, bool) {
	var ok bool
	var token *Token
	var expression abs.ExpressionLike
	var procedure abs.ProcedureLike
	var block abs.BlockLike
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$onClause")
		panic(message)
	}
	_, token, ok = v.parseKeyword("do")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("do",
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$onClause")
		panic(message)
	}
	procedure, token, ok = v.parseProcedure()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$procedure",
			"$procedure",
			"$statements",
			"$statement")
		panic(message)
	}
	block = pro.Block(expression, procedure)
	return block, token, true
}

// This method adds the canonical format for the specified accept clause to the
// state of the formatter.
func (v *formatter) formatBlock(block abs.BlockLike) {
	var expression = block.GetExpression()
	v.formatExpression(expression)
	v.state.AppendString(" do ")
	var procedure = block.GetProcedure()
	v.formatProcedure(procedure)
}

// This method attempts to parse a break clause. It returns the break
// clause and whether or not the break clause was successfully parsed.
func (v *parser) parseBreakClause() (abs.BreakClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var clause abs.BreakClauseLike
	_, token, ok = v.parseKeyword("break")
	if !ok {
		// This is not a break clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("loop")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("loop",
			"$breakClause")
		panic(message)
	}
	clause = pro.BreakClause()
	return clause, token, true
}

// This method adds the canonical format for the specified break clause to the
// state of the formatter.
func (v *formatter) formatBreakClause(clause abs.BreakClauseLike) {
	v.state.AppendString("break loop")
}

// This method attempts to parse a checkout clause. It returns the checkout
// clause and whether or not the checkout clause was successfully parsed.
func (v *parser) parseCheckoutClause() (abs.CheckoutClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient abs.RecipientLike
	var level abs.ExpressionLike
	var moniker abs.ExpressionLike
	var clause abs.CheckoutClauseLike
	_, token, ok = v.parseKeyword("checkout")
	if !ok {
		// This is not a checkout clause.
		return clause, token, false
	}
	recipient, token, ok = v.parseRecipient()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$recipient",
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$ordinal",
			"$moniker")
		panic(message)
	}
	_, token, ok = v.parseKeyword("at")
	if ok {
		// There is an at level part to this clause.
		_, token, ok = v.parseKeyword("level")
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("level",
				"$checkoutClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices",
				"$ordinal",
				"$moniker")
			panic(message)
		}
		level, token, ok = v.parseExpression()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$expression",
				"$checkoutClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices",
				"$ordinal",
				"$moniker")
			panic(message)
		}
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("from",
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$ordinal",
			"$moniker")
		panic(message)
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$moniker",
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$ordinal",
			"$moniker")
		panic(message)
	}
	clause = pro.CheckoutClause(recipient, level, moniker)
	return clause, token, true
}

// This method adds the canonical format for the specified checkout clause to the
// state of the formatter.
func (v *formatter) formatCheckoutClause(clause abs.CheckoutClauseLike) {
	v.state.AppendString("checkout ")
	var recipient = clause.GetRecipient()
	v.formatRecipient(recipient)
	var level = clause.GetLevel()
	if level != nil {
		v.state.AppendString(" at level ")
		v.formatExpression(level)
	}
	v.state.AppendString(" from ")
	var moniker = clause.GetMoniker()
	v.formatExpression(moniker)
}

// This method attempts to parse a continue clause. It returns the continue
// clause and whether or not the continue clause was successfully parsed.
func (v *parser) parseContinueClause() (abs.ContinueClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var clause abs.ContinueClauseLike
	_, token, ok = v.parseKeyword("continue")
	if !ok {
		// This is not a continue clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("loop")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("loop",
			"$continueClause")
		panic(message)
	}
	return clause, token, true
}

// This method adds the canonical format for the specified continue clause to the
// state of the formatter.
func (v *formatter) formatContinueClause(clause abs.ContinueClauseLike) {
	v.state.AppendString("continue loop")
}

// This method attempts to parse a discard clause. It returns the discard
// clause and whether or not the discard clause was successfully parsed.
func (v *parser) parseDiscardClause() (abs.DiscardClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var citation abs.ExpressionLike
	var clause abs.DiscardClauseLike
	_, token, ok = v.parseKeyword("discard")
	if !ok {
		// This is not a discard clause.
		return clause, token, false
	}
	citation, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$document",
			"$discardClause",
			"$document")
		panic(message)
	}
	clause = pro.DiscardClause(citation)
	return clause, token, true
}

// This method adds the canonical format for the specified discard clause to the
// state of the formatter.
func (v *formatter) formatDiscardClause(clause abs.DiscardClauseLike) {
	v.state.AppendString("discard ")
	var citation = clause.GetCitation()
	v.formatExpression(citation)
}

// This method attempts to parse an evaluate clause. It returns the evaluate
// clause and whether or not the evaluate clause was successfully parsed.
func (v *parser) parseEvaluateClause() (abs.EvaluateClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient abs.RecipientLike
	var operator abs.Operator
	var expression abs.ExpressionLike
	var clause abs.EvaluateClauseLike
	recipient, token, ok = v.parseRecipient() // The recipient is optional.
	if ok {
		// The recipient requires an operator.
		operator, token, ok = v.parseOperator()
		if !ok || operator < abs.ASSIGN || operator > abs.QUOTIENT {
			var message = v.formatError(token)
			message += generateGrammar("operator",
				"$evaluateClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices")
			panic(message)
		}
	}
	expression, token, ok = v.parseExpression()
	if !ok {
		// This is not an evaluate clause.
		return clause, token, false
	}
	clause = pro.EvaluateClauseWithRecipient(recipient, operator, expression)
	return clause, token, true
}

// This method adds the canonical format for the specified evaluate clause to the
// state of the formatter.
func (v *formatter) formatEvaluateClause(clause abs.EvaluateClauseLike) {
	if clause.HasRecipient() {
		var recipient, operator = clause.GetRecipient()
		v.formatRecipient(recipient)
		v.state.AppendString(" ")
		v.formatOperator(operator)
		v.state.AppendString(" ")
	}
	var expression = clause.GetExpression()
	v.formatExpression(expression)
}

// This method attempts to parse an if clause. It returns the if clause and
// whether or not the if clause was successfully parsed.
func (v *parser) parseIfClause() (abs.IfClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var block abs.BlockLike
	var clause abs.IfClauseLike
	_, token, ok = v.parseKeyword("if")
	if !ok {
		// This is not an if clause.
		return clause, token, false
	}
	block, token, ok = v.parseBlock()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$condition",
			"$ifClause",
			"$condition",
			"$procedure")
		panic(message)
	}
	clause = pro.IfClause(block)
	return clause, token, true
}

// This method adds the canonical format for the specified if clause to the
// state of the formatter.
func (v *formatter) formatIfClause(clause abs.IfClauseLike) {
	v.state.AppendString("if ")
	var block = clause.GetBlock()
	v.formatBlock(block)
}

// This method attempts to parse a sequence of indices. It returns a list of
// the indices and whether or not the indices were successfully parsed.
func (v *parser) parseIndices() (abs.ListLike[abs.ExpressionLike], *Token, bool) {
	var ok bool
	var token *Token
	var index abs.ExpressionLike
	var indices = col.List[abs.ExpressionLike]()
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not a list of indices.
		return indices, token, false
	}
	index, token, ok = v.parseExpression()
	// There must be at least one index.
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$indices",
			"$expression")
		panic(message)
	}
	for {
		indices.AddValue(index)
		// Every subsequent index must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more indices.
			break
		}
		index, token, ok = v.parseExpression()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$expression",
				"$indices",
				"$expression")
			panic(message)
		}
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("]",
			"$indices",
			"$expression")
		panic(message)
	}
	return indices, token, true
}

// This method adds the canonical format for the specified indices to the
// state of the formatter.
func (v *formatter) formatIndices(indices abs.Sequential[abs.ExpressionLike]) {
	v.state.AppendString("[")
	var iterator = age.Iterator(indices)
	var index = iterator.GetNext() // There is always at least one index.
	v.formatExpression(index)
	for iterator.HasNext() {
		v.state.AppendString(", ")
		index = iterator.GetNext()
		v.formatExpression(index)
	}
	v.state.AppendString("]")
}

// This method attempts to parse a list of inline statements. It returns
// the list of statements and whether or not the list of statements was
// successfully parsed.
func (v *parser) parseInlineStatements() (abs.ListLike[abs.StatementLike], *Token, bool) {
	var ok bool
	var token *Token
	var statement abs.StatementLike
	var statements = col.List[abs.StatementLike]()
	_, token, ok = v.parseDelimiter("}")
	if ok {
		// This is an empty list of statements.
		v.backupOne() // Put back the '}' token.
		return statements, token, true
	}
	statement, token, ok = v.parseStatement()
	if !ok {
		// A non-empty list must have at least one statement.
		var message = v.formatError(token)
		message += generateGrammar("$statement",
			"$statements",
			"$statement")
		panic(message)
	}
	for {
		statements.AddValue(statement)
		// Every subsequent statement must be preceded by a ';'.
		_, token, ok = v.parseDelimiter(";")
		if !ok {
			// No more statements.
			break
		}
		statement, token, ok = v.parseStatement()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$statement",
				"$statements",
				"$statement")
			panic(message)
		}
	}
	return statements, token, true
}

// This method attempts to parse a main clause. It returns the main clause and
// whether or not the main clause was successfully parsed.
func (v *parser) parseMainClause() (abs.ClauseLike, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var mainClause abs.ClauseLike
	mainClause, token, ok = v.parseIfClause()
	if !ok {
		mainClause, token, ok = v.parseSelectClause()
	}
	if !ok {
		mainClause, token, ok = v.parseWithClause()
	}
	if !ok {
		mainClause, token, ok = v.parseWhileClause()
	}
	if !ok {
		mainClause, token, ok = v.parseContinueClause()
	}
	if !ok {
		mainClause, token, ok = v.parseBreakClause()
	}
	if !ok {
		mainClause, token, ok = v.parseReturnClause()
	}
	if !ok {
		mainClause, token, ok = v.parseThrowClause()
	}
	if !ok {
		mainClause, token, ok = v.parseSaveClause()
	}
	if !ok {
		mainClause, token, ok = v.parseDiscardClause()
	}
	if !ok {
		mainClause, token, ok = v.parseNotarizeClause()
	}
	if !ok {
		mainClause, token, ok = v.parseCheckoutClause()
	}
	if !ok {
		mainClause, token, ok = v.parsePublishClause()
	}
	if !ok {
		mainClause, token, ok = v.parsePostClause()
	}
	if !ok {
		mainClause, token, ok = v.parseRetrieveClause()
	}
	if !ok {
		mainClause, token, ok = v.parseAcceptClause()
	}
	if !ok {
		mainClause, token, ok = v.parseRejectClause()
	}
	if !ok {
		// This clause should be checked last since it is slower to fail.
		mainClause, token, ok = v.parseEvaluateClause()
	}
	return mainClause, token, ok
}

// This method adds the canonical format for the specified main clause to the
// state of the formatter.
func (v *formatter) formatMainClause(mainClause abs.ClauseLike) {
	// NOTE: A type switch will only work if each case specifies a unique
	// interface. If two different interfaces define the same method signatures
	// they are indistinguishable as types. To get around this we have added as
	// necessary a unique "dummy" method to each interface to guarantee that it
	// is unique.
	switch value := mainClause.(type) {
	case abs.AcceptClauseLike:
		v.formatAcceptClause(value)
	case abs.BreakClauseLike:
		v.formatBreakClause(value)
	case abs.CheckoutClauseLike:
		v.formatCheckoutClause(value)
	case abs.ContinueClauseLike:
		v.formatContinueClause(value)
	case abs.DiscardClauseLike:
		v.formatDiscardClause(value)
	case abs.EvaluateClauseLike:
		v.formatEvaluateClause(value)
	case abs.IfClauseLike:
		v.formatIfClause(value)
	case abs.NotarizeClauseLike:
		v.formatNotarizeClause(value)
	case abs.PostClauseLike:
		v.formatPostClause(value)
	case abs.PublishClauseLike:
		v.formatPublishClause(value)
	case abs.RejectClauseLike:
		v.formatRejectClause(value)
	case abs.RetrieveClauseLike:
		v.formatRetrieveClause(value)
	case abs.ReturnClauseLike:
		v.formatReturnClause(value)
	case abs.SaveClauseLike:
		v.formatSaveClause(value)
	case abs.SelectClauseLike:
		v.formatSelectClause(value)
	case abs.ThrowClauseLike:
		v.formatThrowClause(value)
	case abs.WhileClauseLike:
		v.formatWhileClause(value)
	case abs.WithClauseLike:
		v.formatWithClause(value)
	default:
		panic(fmt.Sprintf("An invalid main clause (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a list of multiline statements. It returns the
// list of statements and whether or not the list of statements was successfully
// parsed.
func (v *parser) parseMultilineStatements() (abs.ListLike[abs.StatementLike], *Token, bool) {
	var ok bool
	var token *Token
	var statement abs.StatementLike
	var statements = col.List[abs.StatementLike]()
	statement, token, ok = v.parseStatement()
	if !ok {
		// A non-empty list must have at least one statement.
		var message = v.formatError(token)
		message += generateGrammar("$statement",
			"$statements",
			"$statement")
		panic(message)
	}
	for {
		statements.AddValue(statement)
		// Every statement must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$statements",
				"$statement")
			panic(message)
		}
		statement, token, ok = v.parseStatement()
		if !ok {
			// No more statements.
			break
		}
	}
	return statements, token, true
}

// This method attempts to parse a notarize clause. It returns the notarize
// clause and whether or not the notarize clause was successfully parsed.
func (v *parser) parseNotarizeClause() (abs.NotarizeClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var document abs.ExpressionLike
	var moniker abs.ExpressionLike
	var clause abs.NotarizeClauseLike
	_, token, ok = v.parseKeyword("notarize")
	if !ok {
		// This is not a notarize clause.
		return clause, token, false
	}
	document, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$document",
			"$notarizeClause",
			"$document",
			"$moniker")
		panic(message)
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("as",
			"$notarizeClause",
			"$document",
			"$moniker")
		panic(message)
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$moniker",
			"$notarizeClause",
			"$document",
			"$moniker")
		panic(message)
	}
	clause = pro.NotarizeClause(document, moniker)
	return clause, token, true
}

// This method adds the canonical format for the specified notarize clause to the
// state of the formatter.
func (v *formatter) formatNotarizeClause(clause abs.NotarizeClauseLike) {
	v.state.AppendString("notarize ")
	var document = clause.GetDocument()
	v.formatExpression(document)
	v.state.AppendString(" as ")
	var moniker = clause.GetMoniker()
	v.formatExpression(moniker)
}

// This method attempts to parse an exception clause. It returns the exception
// clause and whether or not the exception clause was successfully parsed.
func (v *parser) parseOnClause() (abs.OnClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var exception ele.Symbol
	var block abs.BlockLike
	var blocks = col.List[abs.BlockLike]()
	var clause abs.OnClauseLike
	_, token, ok = v.parseKeyword("on")
	if !ok {
		// This is not an exception clause.
		return clause, token, false
	}
	exception, token, ok = v.parseSymbol()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$exception",
			"$onClause",
			"$exception",
			"$pattern",
			"$statements")
		panic(message)
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		block, token, ok = v.parseBlock()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$pattern",
				"$onClause",
				"$exception",
				"$pattern",
				"$procedure")
			panic(message)
		}
		blocks.AddValue(block)
	}
	// There must be at least one matching block expression.
	if blocks.IsEmpty() {
		var message = v.formatError(token)
		message += generateGrammar("$pattern",
			"$onClause",
			"$exception",
			"$pattern",
			"$procedure")
		panic(message)
	}
	clause = pro.OnClause(exception, blocks)
	return clause, token, true
}

// This method adds the canonical format for the specified on clause to the
// state of the formatter.
func (v *formatter) formatOnClause(clause abs.OnClauseLike) {
	v.state.AppendString("on ")
	var exception = clause.GetException()
	var identifier = exception.GetIdentifier()
	v.state.AppendString("$")
	v.formatIdentifier(identifier)
	var blocks = clause.GetBlocks()
	var iterator = age.Iterator(blocks)
	for iterator.HasNext() {
		var block = iterator.GetNext()
		v.state.AppendString(" matching ")
		v.formatBlock(block)
	}
}

// This method attempts to parse a post clause. It returns the post
// clause and whether or not the post clause was successfully parsed.
func (v *parser) parsePostClause() (abs.PostClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message abs.ExpressionLike
	var bag abs.ExpressionLike
	var clause abs.PostClauseLike
	_, token, ok = v.parseKeyword("post")
	if !ok {
		// This is not a post clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$message",
			"$postClause",
			"$message",
			"$bag")
		panic(message)
	}
	_, token, ok = v.parseKeyword("to")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("to",
			"$postClause",
			"$message",
			"$bag")
		panic(message)
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$bag",
			"$postClause",
			"$message",
			"$bag")
		panic(message)
	}
	clause = pro.PostClause(message, bag)
	return clause, token, true
}

// This method adds the canonical format for the specified post clause to the
// state of the formatter.
func (v *formatter) formatPostClause(clause abs.PostClauseLike) {
	v.state.AppendString("post ")
	var message = clause.GetMessage()
	v.formatExpression(message)
	v.state.AppendString(" to ")
	var bag = clause.GetBag()
	v.formatExpression(bag)
}

// This method attempts to parse a list of statements. It returns the
// list of statements and whether or not the list of statements was
// successfully parsed.
func (v *parser) parseProcedure() (abs.ProcedureLike, *Token, bool) {
	var ok bool
	var token *Token
	var procedure abs.ProcedureLike
	var statements = col.List[abs.StatementLike]()
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		return procedure, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		statements, token, ok = v.parseInlineStatements()
		if !ok {
			v.backupOne() // Put back the '{' character.
			return procedure, token, false
		}
	} else {
		statements, token, ok = v.parseMultilineStatements()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '{' character.
			return procedure, token, false
		}
	}
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("}",
			"$procedure",
			"$statements",
			"$statement")
		panic(message)
	}
	procedure = pro.Procedure(statements)
	return procedure, token, ok
}

// This method adds the canonical format for the specified procedure to the
// state of the formatter.
func (v *formatter) formatProcedure(procedure abs.Sequential[abs.StatementLike]) {
	v.state.AppendString("{")
	switch procedure.GetSize() {
	case 0:
		v.state.AppendString(" ")
	default:
		var iterator = age.Iterator(procedure)
		v.state.IncrementDepth()
		for iterator.HasNext() {
			v.state.AppendNewline()
			var statement = iterator.GetNext()
			v.formatStatement(statement)
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	}
	v.state.AppendString("}")
}

// This method attempts to parse a publish clause. It returns the publish
// clause and whether or not the publish clause was successfully parsed.
func (v *parser) parsePublishClause() (abs.PublishClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var event abs.ExpressionLike
	var clause abs.PublishClauseLike
	_, token, ok = v.parseKeyword("publish")
	if !ok {
		// This is not a publish clause.
		return clause, token, false
	}
	event, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$event",
			"$publishClause",
			"$event")
		panic(message)
	}
	clause = pro.PublishClause(event)
	return clause, token, true
}

// This method adds the canonical format for the specified publish clause to the
// state of the formatter.
func (v *formatter) formatPublishClause(clause abs.PublishClauseLike) {
	v.state.AppendString("publish ")
	var event = clause.GetEvent()
	v.formatExpression(event)
}

// This method attempts to parse a recipient. It returns the recipient and
// whether or not the recipient was successfully parsed.
func (v *parser) parseRecipient() (abs.RecipientLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient abs.RecipientLike
	recipient, token, ok = v.parseSymbol()
	if !ok {
		recipient, token, ok = v.parseAttribute()
	}
	return recipient, token, ok
}

// This method adds the canonical format for the specified recipient to the
// state of the formatter.
func (v *formatter) formatRecipient(recipient abs.RecipientLike) {
	switch value := recipient.(type) {
	case ele.Symbol:
		v.formatSymbol(value)
	case abs.AttributeLike:
		v.formatAttribute(value)
	default:
		panic(fmt.Sprintf("An invalid recipient (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a reject clause. It returns the reject
// clause and whether or not the reject clause was successfully parsed.
func (v *parser) parseRejectClause() (abs.RejectClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message abs.ExpressionLike
	var clause abs.RejectClauseLike
	_, token, ok = v.parseKeyword("reject")
	if !ok {
		// This is not a reject clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$message",
			"$rejectClause",
			"$message")
		panic(message)
	}
	clause = pro.RejectClause(message)
	return clause, token, true
}

// This method adds the canonical format for the specified reject clause to the
// state of the formatter.
func (v *formatter) formatRejectClause(clause abs.RejectClauseLike) {
	v.state.AppendString("reject ")
	var message = clause.GetMessage()
	v.formatExpression(message)
}

// This method attempts to parse a retrieve clause. It returns the retrieve
// clause and whether or not the retrieve clause was successfully parsed.
func (v *parser) parseRetrieveClause() (abs.RetrieveClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient abs.RecipientLike
	var bag abs.ExpressionLike
	var clause abs.RetrieveClauseLike
	_, token, ok = v.parseKeyword("retrieve")
	if !ok {
		// This is not a retrieve clause.
		return clause, token, false
	}
	recipient, token, ok = v.parseRecipient()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$recipient",
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$bag")
		panic(message)
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("from",
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$bag")
		panic(message)
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$bag",
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices",
			"$bag")
		panic(message)
	}
	clause = pro.RetrieveClause(recipient, bag)
	return clause, token, true
}

// This method adds the canonical format for the specified retrieve clause to the
// state of the formatter.
func (v *formatter) formatRetrieveClause(clause abs.RetrieveClauseLike) {
	v.state.AppendString("retrieve ")
	var recipient = clause.GetRecipient()
	v.formatRecipient(recipient)
	v.state.AppendString(" from ")
	var bag = clause.GetBag()
	v.formatExpression(bag)
}

// This method attempts to parse a return clause. It returns the return
// clause and whether or not the return clause was successfully parsed.
func (v *parser) parseReturnClause() (abs.ReturnClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var result abs.ExpressionLike
	var clause abs.ReturnClauseLike
	_, token, ok = v.parseKeyword("return")
	if !ok {
		// This is not a return clause.
		return clause, token, false
	}
	result, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$result",
			"$returnClause",
			"$result")
		panic(message)
	}
	clause = pro.ReturnClause(result)
	return clause, token, true
}

// This method adds the canonical format for the specified return clause to the
// state of the formatter.
func (v *formatter) formatReturnClause(clause abs.ReturnClauseLike) {
	v.state.AppendString("return ")
	var result = clause.GetResult()
	v.formatExpression(result)
}

// This method attempts to parse a save clause. It returns the save
// clause and whether or not the save clause was successfully parsed.
func (v *parser) parseSaveClause() (abs.SaveClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var document abs.ExpressionLike
	var recipient abs.RecipientLike
	var clause abs.SaveClauseLike
	_, token, ok = v.parseKeyword("save")
	if !ok {
		// This is not a save clause.
		return clause, token, false
	}
	document, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$document",
			"$saveClause",
			"$document",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("as",
			"$saveClause",
			"$document",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	recipient, token, ok = v.parseRecipient()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$recipient",
			"$saveClause",
			"$document",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	clause = pro.SaveClause(document, recipient)
	return clause, token, true
}

// This method adds the canonical format for the specified save clause to the
// state of the formatter.
func (v *formatter) formatSaveClause(clause abs.SaveClauseLike) {
	v.state.AppendString("save ")
	var document = clause.GetDocument()
	v.formatExpression(document)
	v.state.AppendString(" as ")
	var recipient = clause.GetRecipient()
	v.formatRecipient(recipient)
}

// This method attempts to parse an select clause. It returns the select clause
// and whether or not the select clause was successfully parsed.
func (v *parser) parseSelectClause() (abs.SelectClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var target abs.ExpressionLike
	var block abs.BlockLike
	var blocks = col.List[abs.BlockLike]()
	var clause abs.SelectClauseLike
	_, token, ok = v.parseKeyword("select")
	if !ok {
		// This is not a select clause.
		return clause, token, false
	}
	target, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$target",
			"$selectClause",
			"$target",
			"$pattern")
		panic(message)
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		block, token, ok = v.parseBlock()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$pattern",
				"$selectClause",
				"$target",
				"$pattern",
				"$procedure")
			panic(message)
		}
		blocks.AddValue(block)
	}
	// There must be at least one matching block expression.
	if blocks.IsEmpty() {
		var message = v.formatError(token)
		message += generateGrammar("$pattern",
			"$selectClause",
			"$target",
			"$pattern")
		panic(message)
	}
	clause = pro.SelectClause(target, blocks)
	return clause, token, true
}

// This method adds the canonical format for the specified select clause to the
// state of the formatter.
func (v *formatter) formatSelectClause(clause abs.SelectClauseLike) {
	v.state.AppendString("select ")
	var target = clause.GetTarget()
	v.formatExpression(target)
	var blocks = clause.GetBlocks()
	var iterator = age.Iterator(blocks)
	for iterator.HasNext() {
		var block = iterator.GetNext()
		v.state.AppendString(" matching ")
		v.formatBlock(block)
	}
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (abs.StatementLike, *Token, bool) {
	var ok bool
	var token *Token
	var statement abs.StatementLike
	var mainClause abs.ClauseLike
	var onClause abs.OnClauseLike
	mainClause, token, ok = v.parseMainClause()
	if ok {
		// The exception clause is optional.
		onClause, token, _ = v.parseOnClause()
	}
	statement = pro.StatementWithHandler(mainClause, onClause)
	return statement, token, ok
}

// This method adds the canonical format for the specified statement to the
// state of the formatter.
func (v *formatter) formatStatement(statement abs.StatementLike) {
	var annotation = statement.GetAnnotation()
	if annotation != nil {
		v.formatAnnotation(annotation)
		v.state.AppendNewline()
	}
	var mainClause = statement.GetMainClause()
	if mainClause != nil {
		v.formatMainClause(mainClause)
	}
	var onClause = statement.GetOnClause()
	if onClause != nil {
		v.formatOnClause(onClause)
	}
	var note = statement.GetNote()
	if note != nil {
		v.formatNote(note)
	}
}

// This method attempts to parse a throw clause. It returns the throw
// clause and whether or not the throw clause was successfully parsed.
func (v *parser) parseThrowClause() (abs.ThrowClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var exception abs.ExpressionLike
	var clause abs.ThrowClauseLike
	_, token, ok = v.parseKeyword("throw")
	if !ok {
		// This is not a throw clause.
		return clause, token, false
	}
	exception, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$exception",
			"$throwClause",
			"$exception")
		panic(message)
	}
	clause = pro.ThrowClause(exception)
	return clause, token, true
}

// This method adds the canonical format for the specified throw clause to the
// state of the formatter.
func (v *formatter) formatThrowClause(clause abs.ThrowClauseLike) {
	v.state.AppendString("throw ")
	var exception = clause.GetException()
	v.formatExpression(exception)
}

// This method attempts to parse a while clause. It returns the while clause and
// whether or not the while clause was successfully parsed.
func (v *parser) parseWhileClause() (abs.WhileClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var block abs.BlockLike
	var clause abs.WhileClauseLike
	_, token, ok = v.parseKeyword("while")
	if !ok {
		// This is not a while clause.
		return clause, token, false
	}
	block, token, ok = v.parseBlock()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$condition",
			"$whileClause",
			"$condition",
			"$procedure")
		panic(message)
	}
	clause = pro.WhileClause(block)
	return clause, token, true
}

// This method adds the canonical format for the specified while clause to the
// state of the formatter.
func (v *formatter) formatWhileClause(clause abs.WhileClauseLike) {
	v.state.AppendString("while ")
	var block = clause.GetBlock()
	v.formatBlock(block)
}

// This method attempts to parse a with clause. It returns the with clause and
// whether or not the with clause was successfully parsed.
func (v *parser) parseWithClause() (abs.WithClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var value ele.Symbol
	var block abs.BlockLike
	var clause abs.WithClauseLike
	_, token, ok = v.parseKeyword("with")
	if !ok {
		// This is not a with clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("each")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("each",
			"$withClause",
			"$value",
			"$sequence")
		panic(message)
	}
	value, token, ok = v.parseSymbol()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$value",
			"$withClause",
			"$value",
			"$sequence")
		panic(message)
	}
	_, token, ok = v.parseKeyword("in")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("in",
			"$withClause",
			"$value",
			"$sequence")
		panic(message)
	}
	block, token, ok = v.parseBlock()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$sequence",
			"$withClause",
			"$value",
			"$sequence",
			"$procedure")
		panic(message)
	}
	clause = pro.WithClause(value, block)
	return clause, token, true
}

// This method adds the canonical format for the specified with clause to the
// state of the formatter.
func (v *formatter) formatWithClause(clause abs.WithClauseLike) {
	v.state.AppendString("with each ")
	var value = clause.GetValue()
	var identifier = value.GetIdentifier()
	v.state.AppendString("$")
	v.formatIdentifier(identifier)
	v.state.AppendString(" in ")
	var block = clause.GetBlock()
	v.formatBlock(block)
}
