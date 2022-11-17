package tokens

import "analizador-lexico/lexer_framework"

var Operadores = []string{
	"OPERADOR_ARITMETICO",
	"OPERADOR_RELACIONAL",
	"OPERADOR_LOGICO",
	"OPERADOR_ASIGNACION",
	"OPERADOR_TERMINAL",
	"ULTIMO_INDICE",
}

func AgregarOperadores() {
	lexer_framework.Lexer.Add([]byte(`[|&~]`), lexer_framework.BuildToken("OPERADOR_LOGICO"))
	lexer_framework.Lexer.Add([]byte(`[=]`), lexer_framework.BuildToken("OPERADOR_ASIGNACION"))
	lexer_framework.Lexer.Add([]byte(`(<=|<|>|>=|==|~=)`), lexer_framework.BuildToken("OPERADOR_RELACIONAL"))
	lexer_framework.Lexer.Add([]byte(`[+\-*\\\/\^]`), lexer_framework.BuildToken("OPERADOR_ARITMETICO"))
	lexer_framework.Lexer.Add([]byte(`[;]`), lexer_framework.BuildToken("OPERADOR_TERMINAL"))
	lexer_framework.Lexer.Add([]byte(`[$]`), lexer_framework.BuildToken("ULTIMO_INDICE"))
}
