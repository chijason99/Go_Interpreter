package lexer

import (
	"testing"

	"chijason99/northwet_interpreter/token"
)

func Test_Token(t *testing.T) {
	input := "+;{}="

	testCases := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{ token.PLUS, "+" },
		{ token.SEMICOLON, ";"},
		{ token.LBRACE, "{"},
		{ token.RBRACE, "}"},
		{ token.ASSIGN, "="},
		{ token.EOF, ""},
	}

	l := New(input)

	for i, tc := range testCases {
		tok := l.NextToken()

		if tc.expectedType != tok.Type {
			t.Fatalf("Tests %d: Expected %q but got %q", i, tc.expectedType, tok.Type)
		}

		if tc.expectedLiteral != tok.Literal {
			t.Fatalf("Tests %d: Expected %q but got %q", i, tc.expectedLiteral, tok.Literal)
		}
	}
}

func Test_Next_Token(t *testing.T){
	input := `let five = 5;
		let ten = 10;

		let add = func(x, y){
			return x + y;
		};

		!-/*5;

		5 < 10 > 5;

		if (5 < 10){
			return true;
		} else {
			return false; 
		}

		let result = add(five, ten);

		10 == 10;
		9 != 10;
		9 <= 10;
		10 >= 9;
	`

	testCases := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{ token.LET, "let" },
		{ token.IDENT, "five"},
		{ token.ASSIGN, "="},
		{ token.INT, "5"},
		{ token.SEMICOLON, ";"},

		{ token.LET, "let"},
		{ token.IDENT, "ten" },
		{ token.ASSIGN, "="},
		{ token.INT, "10"},
		{ token.SEMICOLON, ";"},

		{ token.LET, "let"},
		{ token.IDENT, "add"},
		{ token.ASSIGN, "=" },
		{ token.FUNCTION, "func"},
		{ token.LPAREN, "("},
		{ token.IDENT, "x"},
		{ token.COMMA, ","},
		{ token.IDENT, "y" },
		{ token.RPAREN, ")"},
		{ token.LBRACE, "{"},
		{ token.RETURN, "return"},
		{ token.IDENT, "x"},
		{ token.PLUS, "+"},
		{ token.IDENT, "y"},		
		{ token.SEMICOLON, ";"},		
		{ token.RBRACE, "}"},
		{ token.SEMICOLON, ";"},

		{ token.NEGATION, "!"},		
		{ token.MINUS, "-"},		
		{ token.SLASH, "/"},		
		{ token.ASTERISK, "*"},		
		{ token.INT, "5"},		
		{ token.SEMICOLON, ";"},

		{ token.INT, "5"},
		{ token.LT, "<"},
		{ token.INT, "10"},	
		{ token.GT, ">"},	
		{ token.INT, "5"},
		{ token.SEMICOLON, ";"},

		{ token.IF, "if"},
		{ token.LPAREN, "("},		
		{ token.INT, "5"},
		{ token.LT, "<"},		
		{ token.INT, "10"},
		{ token.RPAREN, ")"},
		{ token.LBRACE, "{"},
		{ token.RETURN, "return"},
		{ token.TRUE, "true"},
		{ token.SEMICOLON, ";"},
		{ token.RBRACE, "}"},
		{ token.ELSE, "else"},		
		{ token.LBRACE, "{"},
		{ token.RETURN, "return"},		
		{ token.FALSE, "false"},
		{ token.SEMICOLON, ";"},
		{ token.RBRACE, "}"},

		{ token.LET, "let"},
		{ token.IDENT, "result"},		
		{ token.ASSIGN, "="},
		{ token.IDENT, "add"},		
		{ token.LPAREN, "("},
		{ token.IDENT, "five"},
		{ token.COMMA, ","},
		{ token.IDENT, "ten"},
		{ token.RPAREN, ")"},
		{ token.SEMICOLON, ";"},

		{ token.INT, "10"},
		{ token.EQ, "=="},
		{ token.INT, "10"},
		{ token.SEMICOLON, ";"},

		{ token.INT, "9"},
		{ token.NOT_EQ, "!="},
		{ token.INT, "10"},
		{ token.SEMICOLON, ";"},

		{ token.INT, "9"},
		{ token.LE, "<="},
		{ token.INT, "10"},
		{ token.SEMICOLON, ";"},

		{ token.INT, "10"},
		{ token.GE, ">="},
		{ token.INT, "9"},
		{ token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tc := range testCases {
		tok := l.NextToken()

		if tc.expectedLiteral != tok.Literal {
			t.Fatalf("Tests %d: Expected %q but got %q", i, tc.expectedLiteral, tok.Literal)
		}

		if tc.expectedType != tok.Type {
			t.Fatalf("Tests %d: Expected %q but got %q", i, tc.expectedType, tok.Type)
		}
	}
}