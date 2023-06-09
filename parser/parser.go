package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = loadPrefixesFuncs(p)
	p.infixParseFns = loadInfixFuncs(p)

	// Read two tokens, so curToen and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()

	for p.curToken.Type != token.Eof {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !expectPeek(p, token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !expectPeek(p, token.Assing) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(Lowest)

	if peekTokenIs(p, token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(Lowest)

	if peekTokenIs(p, token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(Lowest)

	if peekTokenIs(p, token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		noPrefixParseFnError(p, p.curToken.Type)
		return nil
	}

	leftExp := prefix()
	for !peekTokenIs(p, token.Semicolon) && precedence < peekPrecedence(p) {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(Prefix)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedenc := curPrecedence(p)
	p.nextToken()
	expression.Right = p.parseExpression(precedenc)

	return expression
}

func (p *Parser) parseBooleanExpression() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: curTokenIs(p, token.True)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	ext := p.parseExpression(Lowest)

	if !expectPeek(p, token.Rparen) {
		return nil
	}

	return ext
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	if !expectPeek(p, token.Lparen) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(Lowest)

	if !expectPeek(p, token.Rparen) {
		return nil
	}

	if !expectPeek(p, token.Lbrace) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if peekTokenIs(p, token.Else) {
		p.nextToken()

		if !expectPeek(p, token.Lbrace) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = make([]ast.Statement, 0)

	p.nextToken()

	for !curTokenIs(p, token.Rbrace) && !curTokenIs(p, token.Eof) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !expectPeek(p, token.Lparen) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !expectPeek(p, token.Lbrace) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := make([]*ast.Identifier, 0)

	if peekTokenIs(p, token.Rparen) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for peekTokenIs(p, token.Comma) {
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !expectPeek(p, token.Rparen) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	return &ast.CallExpression{
		Token:     p.curToken,
		Function:  function,
		Arguments: p.parseExpressionList(token.Rparen),
	}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return ast.NewStringLiteral(p.curToken, p.curToken.Literal)
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := ast.NewArray(p.curToken)
	array.Elements = p.parseExpressionList(token.Rbracket)

	return array
}

func (p *Parser) parseExpressionList(endToken token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if peekTokenIs(p, endToken) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(Lowest))

	for peekTokenIs(p, token.Comma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(Lowest))
	}

	if !expectPeek(p, endToken) {
		return nil
	}

	return list
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(Lowest)

	if !expectPeek(p, token.Rbracket) {
		return nil
	}

	return exp
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := ast.NewHashLiteral(p.curToken)

	for !peekTokenIs(p, token.Rbrace) {
		p.nextToken()
		key := p.parseExpression(Lowest)

		if !expectPeek(p, token.Collon) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(Lowest)

		hash.Pairs[key] = value

		if !peekTokenIs(p, token.Rbrace) && !expectPeek(p, token.Comma) {
			return nil
		}
	}

	if !expectPeek(p, token.Rbrace) {
		return nil
	}

	return hash
}
