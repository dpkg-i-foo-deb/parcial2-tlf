package tokens

import "analizador/lexer_framework"

var Condicionales = []string{
	"CONDICIONAL IF",
	"CONDICIONAL ELSE",
	"CONDICIONAL ELSEIF",
	"PALABRA RESERVADA THEN",
	"PALABRA RESERVADA END",
}

func AgregarCondicionales() {
	lexer_framework.Lexer.Add([]byte(`if`), lexer_framework.BuildToken("CONDICIONAL IF"))
	lexer_framework.Lexer.Add([]byte(`else`), lexer_framework.BuildToken("CONDICIONAL ELSE"))
	lexer_framework.Lexer.Add([]byte(`elseif`), lexer_framework.BuildToken("CONDICIONAL ELSEIF"))
	lexer_framework.Lexer.Add([]byte(`then`), lexer_framework.BuildToken("PALABRA RESERVADA THEN"))
	lexer_framework.Lexer.Add([]byte(`end`), lexer_framework.BuildToken("PALABRA RESERVADA END"))
}
