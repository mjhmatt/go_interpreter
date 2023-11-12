package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"testing"
)

func TestLetStatemenets(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 832832;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatal("program does not have 3 statements has", len(program.Statements))

	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	lenErrors := len(errors)
	if lenErrors == 0 {
		return
	}
	t.Errorf("parser had %d errors", lenErrors)
	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Literal not let got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("name not name")
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("token literal name doesnt match")
		return false
	}
	return true
}
