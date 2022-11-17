package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var OperacionImprimir util.Expresion
var ParametrosImprimir util.Expresion

func InicializarImpresion() {
	exp, err := regexp.Compile(`print`)

	OperacionImprimir = util.Expresion{
		Regex:     exp,
		Categoria: "OPERACION_IMPRIMIR",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`\('([0-9]|[a-z-A-Z]|\.)+',([a-z]|[#%!$?_].[^A-Z]).([a-zA-Z]*|[0-9]*|[#%!$?_]*)*\)`)

	ParametrosImprimir = util.Expresion{
		Regex:     exp,
		Categoria: "PARAMETROS_IMPRIMIR",
	}

	util.VerificarError(err)
}
