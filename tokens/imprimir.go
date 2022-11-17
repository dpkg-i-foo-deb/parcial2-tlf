package tokens

import "analizador/lexer_framework"

var Imprimir = []string{
	"IMPRIMIR",
}

func AgregarImpresiones() {
	lexer_framework.Lexer.Add([]byte(`(print)\(.*\)`), lexer_framework.BuildToken("IMPRIMIR"))

}
