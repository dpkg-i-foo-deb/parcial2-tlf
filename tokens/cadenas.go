package tokens

import "analizador-lexico/lexer_framework"

var Cadenas = []string{
	"CADENA",
}

func AgregarCadenas() {
	lexer_framework.Lexer.Add([]byte(`("|')([a-zA-Z]|[0-9])*("|')`), lexer_framework.BuildToken("CADENA"))
}
