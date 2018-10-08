package lexer

import (
	"testing"

	"github.com/itkq/go-minruby-interpreter/token"
	"github.com/k0kubun/pp"
)

func TestNextToken(t *testing.T) {
	input := `def even?(n)
  if n == 0
    true
  else
    odd?(n - 1)
  end
end

def odd?(n)
  if n == 0
    false
  else
    even?(n - 1)
  end
end

p(even?(10))
p(even?(11))
p(odd?(10))
p(odd?(11))
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "def"},
		{token.IDENT, "even?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.RPAREN, ")"},
		{token.IF, "if"},
		{token.IDENT, "n"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.TRUE, "true"},
		{token.ELSE, "else"},
		{token.IDENT, "odd?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.END, "end"},
		{token.END, "end"},
		{token.FUNCTION, "def"},
		{token.IDENT, "odd?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.RPAREN, ")"},
		{token.IF, "if"},
		{token.IDENT, "n"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.FALSE, "false"},
		{token.ELSE, "else"},
		{token.IDENT, "even?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.END, "end"},
		{token.END, "end"},
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.IDENT, "even?"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.IDENT, "even?"},
		{token.LPAREN, "("},
		{token.INT, "11"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.IDENT, "odd?"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.IDENT, "odd?"},
		{token.LPAREN, "("},
		{token.INT, "11"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestComment(t *testing.T) {
	input := `# This is a comment.
def even?(n)
  # comment
  if n == 0
    true # comment
  else
    odd?(n - 1)
  end
end # comment`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "def"},
		{token.IDENT, "even?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.RPAREN, ")"},
		{token.IF, "if"},
		{token.IDENT, "n"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.TRUE, "true"},
		{token.ELSE, "else"},
		{token.IDENT, "odd?"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.END, "end"},
		{token.END, "end"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		pp.Print(tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
