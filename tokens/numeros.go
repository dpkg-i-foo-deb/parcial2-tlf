package tokens

import "analizador-lexico/lexer_framework"

var Numeros = []string{
	"ENTERO",
	"DECIMAL",
}

func AgregarNumeros() {
	lexer_framework.Lexer.Add([]byte(`([0-9]|-[0-9])[0-9]*`), lexer_framework.BuildToken("ENTERO"))
	lexer_framework.Lexer.Add([]byte(`([0-9]|-[0-9])[0-9]*\.[0-9][0-9]*`), lexer_framework.BuildToken("DECIMAL"))
}
