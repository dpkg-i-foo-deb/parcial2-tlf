package tokens

import "analizador/lexer_framework"

var Condicionales = []string{
	"CONDICIONALES",
}

func AgregarCondicionales() {
	lexer_framework.Lexer.Add([]byte(`(if|else|elseif)`), lexer_framework.BuildToken("CONDICIONALES"))
}
