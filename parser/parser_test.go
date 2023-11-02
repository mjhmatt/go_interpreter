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
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatal("profram does not have 3 statements has", len(program.Statements))
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
