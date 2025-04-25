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
		if l.peekChar() == '=' {
			current := l.ch
			l.readChar()
			
			tokType = token.EQ
			lit = string(current) + string(l.ch)
		} else {
			tokType = token.ASSIGN
		}
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
		if l.peekChar() == '=' {
			current := l.ch
			l.readChar()
			
			tokType = token.NOT_EQ
			lit = string(current) + string(l.ch)
		} else {
			tokType = token.NEGATION
		}
	case '<':
		if l.peekChar() == '=' {
			current := l.ch
			l.readChar()
			
			tokType = token.LE
			lit = string(current) + string(l.ch)
		} else {
			tokType = token.LT
		}
	case '>':
		if l.peekChar() == '=' {
			current := l.ch
			l.readChar()
			
			tokType = token.GE
			lit = string(current) + string(l.ch)
		} else {
			tokType = token.GT
		}
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
			return newToken(tokType, lit)
		} else if isDigit(l.ch) {
			lit = l.readNumber()

			// Return early to avoid readChar later that would skip the current character
			return newToken(token.INT, lit)
		}else {
			tokType = token.ILLEGAL
		}
	}

	tok := newToken(tokType, lit)
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

// Like readChar, the only difference is that it would only peek, without incrementing the position of the lexer
func (l *Lexer) peekChar() byte{
	if (l.readPosition >= len(l.input)){
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokType token.TokenType, lit string) token.Token {
	return token.Token{ Type: tokType, Literal: lit }
}