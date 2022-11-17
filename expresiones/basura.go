package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

// Aquí se colocan las cosas que deseamos ignorar
var SaltoLinea util.Expresion

func InicializarBasura() {
	exp, err := regexp.Compile("\n")

	SaltoLinea = util.Expresion{
		Regex:     exp,
		Categoria: "SALTO_LINEA",
	}

	util.VerificarError(err)
}
