package tokens

import "analizador-lexico/lexer_framework"

var Imprimir = []string{
	"IMPRIMIR",
}

func AgregarImpresiones() {
	lexer_framework.Lexer.Add([]byte(`(print)\('([a-zA-Z]*|[0-9]*|\.*)*'(,([a-zA-z]*|[0-9]*)*)+\)`), lexer_framework.BuildToken("IMPRIMIR"))

}
