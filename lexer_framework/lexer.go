package lexer_framework

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lexmachine.Lexer // The lexer object. Use this to construct a Scanner
var Unrecognized_Tokens []string

func initTokens() {
	Tokens = []string{
		"ID",
		"SEPARADOR_SENTENCIAS",
		"ULTIMO_INDICE",
		"CICLOS",
	}

	Unrecognized_Tokens = []string{
		"NOT_VALID",
	}

}

// Creates the lexer object and compiles the NFA.
func InitLexer() {
	Lexer = lexmachine.NewLexer()
	//Lexer.Add([]byte(`[$]`), BuildToken("ULTIMO_INDICE"))
	//Lexer.Add([]byte("( |\t|\n|\r)+"), skip)
	//Lexer.Add([]byte(`ñ`), unrecognized)
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

func unrecognized(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	return s.Token(TokenIds["NOT_VALID"], string(m.Bytes), m), nil
}

// a lexmachine.Action function with constructs a Token of the given token type by
// the token type's name.
func BuildToken(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}
