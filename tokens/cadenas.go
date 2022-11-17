package tokens

import "analizador/lexer_framework"

var Cadenas = []string{
	"CADENA",
}

func AgregarCadenas() {
	lexer_framework.Lexer.Add([]byte(`[\'\"][a-z]*[\'\"]`), lexer_framework.BuildToken("CADENA"))
}
