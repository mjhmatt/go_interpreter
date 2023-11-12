package parser

import (
	"fmt"
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"go_interpreter/token"
)

// Parser struct to handle parsing
type Parser struct {
	l            *lexer.Lexer // Pointer to the lexer instance
	currentToken token.Token  // Current token being processed
	peekToken    token.Token  // Next token to be processed
	errors       []string
}

// New creates a new Parser with a lexer and advances it to the next token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken() // Read the first token
	p.nextToken() // Read the next token

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expect next token to be %s got %s instead", t, p.peekToken.Type))
}

// nextToken advances the parser to the next token by updating current and peek tokens
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the entire program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Loop through tokens until reaching the end of the file
	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement() // Parse individual statements
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) // Append parsed statement to the program
		}
		p.nextToken() // Move to the next token
	}

	return program
}

// parseStatement parses a single statement based on the current token type
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement() // If it's a 'let' statement, parse it
	default:
		return nil // If it's not recognized, return nil
	}
}

// parseLetStatement parses a 'let' statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken} // Create a new 'let' statement object

	// Check if the next token is an identifier (variable name)
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	// Assign the name of the variable to the 'let' statement
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// Check if the token after the identifier is an assignment token
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken() // Move to the next token after the assignment

	// If the current token is not a semicolon, move to the next token
	if !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// currentTokenIs checks if the current token type matches a given token type
func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

// peekTokenIs checks if the next token (peek token) type matches a given token type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks if the next token matches the given type and advances if it does
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken() // Advance to the next token
		return true
	}
	p.peekError(t)
	return false
}
