package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"go_interpreter/token"
)

// l is a pointer to instance of a leer which we call next token on to get the next token
// peek is to determine if we are EOL or not or at start of an expression
type Parser struct {
	l            *lexer.Lexer //pointer to the lexer instance
	currentToken token.Token  // current token being processed
	peekToken    token.Token  // next token to be processed
}

// create a new Parser with a lexer and advances it to next token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

// advances parser to next token by update current and peek token
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
