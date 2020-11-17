package parser

import (
	"github.com/hiroshimashu/monkey-interpreter/src/lexer"
	"github.com/hiroshimashu/monkey-interpreter/src/token"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l:l}

	p.nextToken()
	p.nextTokne()

	return p 
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}