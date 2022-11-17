package tokens

import "analizador/lexer_framework"

var NombresVariables = []string{
	"NOMBRE_VARIABLE",
}

func AgregarNombresVariables() {
	lexer_framework.Lexer.Add([]byte(`([a-z]|[#%!$?_].[^A-Z]).([a-zA-Z]*|[0-9]*|[#%!$?_]*)*`), lexer_framework.BuildToken("NOMBRE_VARIABLE"))
}
