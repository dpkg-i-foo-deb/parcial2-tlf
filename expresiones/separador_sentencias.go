package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var SeparadorSentencias util.Expresion

func InicializarSeparadorSentencias() {
	exp, err := regexp.Compile(`[,;]`)

	SeparadorSentencias = util.Expresion{
		Regex:     exp,
		Categoria: "SEPARADOR_SENTENCIAS",
	}

	util.VerificarError(err)
}
