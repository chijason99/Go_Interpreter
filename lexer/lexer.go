package lexer

import (
	"chijason99/northwet_interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // The current position
	readPosition int  // The next position to read
	ch           byte // The character being examined right now
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()

	var tokType token.TokenType
	lit := string(l.ch)

	switch l.ch {
	case '=':
		tokType = token.ASSIGN
	case '+':
		tokType = token.PLUS
	case ',':
		tokType = token.COMMA
	case ';':
		tokType = token.SEMICOLON
	case '(':
		tokType = token.LPAREN
	case ')':
		tokType = token.RPAREN
	case '{':
		tokType = token.LBRACE
	case '}':
		tokType = token.RBRACE
	case '!':
		tokType = token.NEGATION
	case '<':
		tokType = token.LT
	case '>':
		tokType = token.GT
	case '/':
		tokType = token.SLASH
	case '*':
		tokType = token.ASTERISK
	case '-':
		tokType = token.MINUS
	case 0:
		tokType = token.EOF
		lit = ""
	default:
		if isLetter(l.ch) {
			lit = l.readIdentifier()
			tokType = token.LookUpIdentifier(lit)

			// Return early to avoid readChar later that would skip the current character
			return token.Token{ Type: tokType, Literal: lit }
		} else if isDigit(l.ch) {
			lit = l.readNumber()
			tokType = token.INT

			// Return early to avoid readChar later that would skip the current character
			return token.Token{ Type: tokType, Literal: lit }
		}else {
			tokType = token.ILLEGAL
		}
	}

	tok := token.Token{Type: tokType, Literal: lit}
	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string{
	position := l.position

	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string{
	position := l.position

	for isDigit(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool{
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && 'Z' >= ch) || ch == '_'
}

func isDigit(ch byte) bool{
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipWhiteSpace(){
	for l.ch == ' ' || l.ch == '\r' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}