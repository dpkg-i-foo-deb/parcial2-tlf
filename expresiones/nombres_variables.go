package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var NombreVariables util.Expresion

func InicializarNombres() {
	exp, err := regexp.Compile(`([a-z]|[#%!$?_].[^A-Z]).([a-zA-Z]*|[0-9]*|[#%!$?_]*)*`)

	NombreVariables = util.Expresion{
		Regex:     exp,
		Categoria: "NOMBRE_VARIABLE",
	}

	util.VerificarError(err)

}
