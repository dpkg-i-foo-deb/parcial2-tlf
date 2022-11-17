package tokens

import "analizador-lexico/lexer_framework"

var Comentarios = []string{
	"COMENTARIO",
}

func AgregarComentarios() {
	lexer_framework.Lexer.Add([]byte(`//[^\n]*\n?`), lexer_framework.BuildToken("COMENTARIO"))
	lexer_framework.Lexer.Add([]byte(`\/\*([^*]|\r|\n|(\*+([^*\/]|\r|\n)))*\*\/`), lexer_framework.BuildToken("COMENTARIO"))
}
