package tokens

import "analizador/lexer_framework"

var SeparadorSentencias = []string{
	"SEPARADOR_SENTENCIAS",
}

func AgregarSeparadorSentencias() {
	lexer_framework.Lexer.Add([]byte(`[,;]`), lexer_framework.BuildToken("SEPARADOR_SENTENCIAS"))
}
