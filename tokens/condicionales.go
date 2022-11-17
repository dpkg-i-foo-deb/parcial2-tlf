package tokens

import "analizador-lexico/lexer_framework"

var Condicionales = []string{
	"SENTENCIA CONDICIONAL",
	"CONDICIONAL ELSE",
	"CONDICIONAL ELSEIF",
	"PALABRA RESERVADA THEN",
	"PALABRA RESERVADA END",
	"PREGUNTA DEL CONDICIONAL",
}

func AgregarCondicionales() {
	lexer_framework.Lexer.Add([]byte(`if`), lexer_framework.BuildToken("SENTENCIA CONDICIONAL"))
	lexer_framework.Lexer.Add([]byte(`else`), lexer_framework.BuildToken("CONDICIONAL ELSE"))
	lexer_framework.Lexer.Add([]byte(`elseif`), lexer_framework.BuildToken("CONDICIONAL ELSEIF"))
	lexer_framework.Lexer.Add([]byte(`then`), lexer_framework.BuildToken("PALABRA RESERVADA THEN"))
	lexer_framework.Lexer.Add([]byte(`end`), lexer_framework.BuildToken("PALABRA RESERVADA END"))
	//lexer_framework.Lexer.Add([]byte(`(\(.([a-zaAZ]*|[0-9]*).\)){1}`), lexer_framework.BuildToken("PREGUNTA DEL CONDICIONAL"))

	lexer_framework.Lexer.Add([]byte(`\(`), lexer_framework.BuildToken("PREGUNTA DEL CONDICIONAL"))

}
