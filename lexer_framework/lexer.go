package lexer_framework

import (
	"strings"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var Operadores []string     // The tokens representing literal strings
var Keywords []string       // The keyword tokens
var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lexmachine.Lexer // The lexer object. Use this to construct a Scanner
var Unrecognized_Tokens []string

// Called at package initialization. Creates the lexer and populates token lists.
func init() {
	initTokens()
	var err error
	Lexer, err = initLexer()
	if err != nil {
		panic(err)
	}
}

func initTokens() {
	Tokens = []string{
		"COMMENTARIO",
		"ID",
		"OPERADOR_ARITMETICO",
		"OPERADOR_RELACIONAL",
		"OPERADOR_LOGICO",
		"OPERADOR_ASIGNACION",
		"SEPARADOR_SENTENCIAS",
		"OPERADOR_TERMINAL",
		"ULTIMO_INDICE",
		"CICLOS",
	}

	Unrecognized_Tokens = []string{
		"NOT_VALID",
	}

	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

// Creates the lexer object and compiles the NFA.
func initLexer() (*lexmachine.Lexer, error) {
	Lexer = lexmachine.NewLexer()

	for _, lit := range Operadores {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		Lexer.Add([]byte(r), token(lit))
	}
	for _, name := range Keywords {
		Lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	Lexer.Add([]byte(`//[^\n]*\n?`), token("COMMENT"))
	Lexer.Add([]byte(`\/\*([^*]|\r|\n|(\*+([^*\/]|\r|\n)))*\*\/`), token("COMMENTARIO"))
	Lexer.Add([]byte(`[+\-*\\\/\^]`), token("OPERADOR_ARITMETICO"))
	Lexer.Add([]byte(`(<=|<|>|>=|==|~=)`), token("OPERADOR_RELACIONAL"))
	Lexer.Add([]byte(`[|&~]`), token("OPERADOR_LOGICO"))
	Lexer.Add([]byte(`[=]`), token("OPERADOR_ASIGNACION"))
	Lexer.Add([]byte(`[,;]`), token("SEPARADOR_SENTENCIAS"))
	Lexer.Add([]byte(`[;]`), token("OPERADOR_TERMINAL"))
	Lexer.Add([]byte(`[$]`), token("ULTIMO_INDICE"))
	Lexer.Add([]byte("( |\t|\n|\r)+"), skip)
	Lexer.Add([]byte(`ñ`), unrecognized)

	err := Lexer.Compile()
	if err != nil {
		return nil, err
	}
	return Lexer, nil
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
func token(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}
