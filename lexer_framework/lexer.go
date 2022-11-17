package lexer_framework

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lexmachine.Lexer // The lexer object. Use this to construct a Scanner
func initTokens() {
	Tokens = []string{}

}

// Creates the lexer object and compiles the NFA.
func InitLexer() {
	Lexer = lexmachine.NewLexer()
	Lexer.Add([]byte("( |\t|\n|\r)+"), skip)
}

func Compilar() error {

	err := Lexer.Compile()
	if err != nil {
		return err
	}
	return nil
}

// a lexmachine.Action function which skips the match.
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lexmachine.Action function with constructs a Token of the given token type by
// the token type's name.
func BuildToken(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}
