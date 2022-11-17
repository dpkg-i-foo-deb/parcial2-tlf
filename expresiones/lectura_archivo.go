package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var LectorArchivo util.Expresion
var ParametrosLectorArchivo util.Expresion

func InicializarLector() {
	exp, err := regexp.Compile(`mgetl`)

	LectorArchivo = util.Expresion{
		Regex:     exp,
		Categoria: "FUNCIÓN LECTURA DE ARCHIVOS",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(``)

	LectorArchivo = util.Expresion{
		Regex:     exp,
		Categoria: "PARÁMETROS FUNCIÓN LECTURA DE ARCHIVOS",
	}

	util.VerificarError(err)
}
