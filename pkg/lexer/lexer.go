package lexer

import "github.com/Amerosa/monkey/pkg/token"

type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0 //NUL in ASCII
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            literal := string(ch) + string(l.ch)
            tok = token.Token{Type: token.EQ, Literal: literal}
        } else {
            tok = token.NewToken(token.ASSIGN, l.ch)
        }
    case '-':
        tok = token.NewToken(token.MINUS, l.ch)
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            literal := string(ch) + string(l.ch)
            tok = token.Token{Type: token.NOT_EQ, Literal: literal}
        } else {
            tok = token.NewToken(token.BANG, l.ch)
        }
    case '/':
        tok = token.NewToken(token.SLASH, l.ch)
    case '*':
        tok = token.NewToken(token.ASTERISK, l.ch)
    case '<':
        tok = token.NewToken(token.LT, l.ch)
    case '>':
        tok = token.NewToken(token.GT, l.ch)
    case ';':
        tok = token.NewToken(token.SEMICOLON, l.ch)
    case '(':
        tok = token.NewToken(token.LPAREN, l.ch)
    case ')':
        tok = token.NewToken(token.RPAREN, l.ch)
    case ',':
        tok = token.NewToken(token.COMMA, l.ch)
    case '+':
        tok = token.NewToken(token.PLUS, l.ch)
    case '{':
        tok = token.NewToken(token.LBRACE, l.ch)
    case '}':
        tok = token.NewToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            ident := l.readIdentifier()
            tok.Type = token.LookupIdent(ident)
            tok.Literal = ident
            return tok
        } else if isDigit(l.ch) {
            tok.Type = token.INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = token.NewToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

func isDigit (ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func isLetter (ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
