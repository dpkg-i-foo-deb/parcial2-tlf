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
	Operadores = []string{}
	Keywords = []string{
		"PUBLIC",
		"STATIC",
		"VOID",
		"MAIN",
		"INT",
		"STRICT",
	}
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
	}

	Unrecognized_Tokens = []string{
		"NOT_VALID",
	}

	Tokens = append(Tokens, Keywords...)
	Tokens = append(Tokens, Operadores...)
	Tokens = append(Tokens, Unrecognized_Tokens...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

// Creates the lexer object and compiles the NFA.
func initLexer() (*lexmachine.Lexer, error) {
	lexer := lexmachine.NewLexer()

	for _, lit := range Operadores {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}
	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	lexer.Add([]byte(`//[^\n]*\n?`), token("COMMENT"))
	lexer.Add([]byte(`\/\*([^*]|\r|\n|(\*+([^*\/]|\r|\n)))*\*\/`), token("COMMENTARIO"))
	lexer.Add([]byte(`[+\-*\\\/\^]`), token("OPERADOR_ARITMETICO"))
	lexer.Add([]byte(`(<=|<|>|>=|==|~=)`), token("OPERADOR_RELACIONAL"))
	lexer.Add([]byte(`[|&~]`), token("OPERADOR_LOGICO"))
	lexer.Add([]byte(`[=]`), token("OPERADOR_ASIGNACION"))
	lexer.Add([]byte(`[,;]`), token("SEPARADOR_SENTENCIAS"))
	lexer.Add([]byte(`[;]`), token("OPERADOR_TERMINAL"))
	lexer.Add([]byte(`[$]`), token("ULTIMO_INDICE"))
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)
	lexer.Add([]byte(`ñ`), unrecognized)

	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
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
