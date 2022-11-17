package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var LectorArchivo util.Expresion
var ParametrosLectorArchivo util.Expresion

func InicializarLectorArchivo() {
	exp, err := regexp.Compile(`mgetl`)

	LectorArchivo = util.Expresion{
		Regex:     exp,
		Categoria: "FUNCIÓN LECTURA DE ARCHIVOS",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`\(((\/?\w+\/)*\w+\.\w+|'(\/?\w+\/)*\w+\.\w+'|\((\/?\w+\/)*\w+\.\w+\)|\('(\/?\w+\/)*\w+\.\w+'\))\)|((\/?\w+\/)*\w+\.\w+|'(\/?\w+\/)*\w+\.\w+'|\((\/?\w+\/)*\w+\.\w+\)|\('(\/?\w+\/)*\w+\.\w+'\))|\(((\/?\w+\/)*\w+\.\w+|'(\/?\w+\/)*\w+\.\w+'|\((\/?\w+\/)*\w+\.\w+\)|\('(\/?\w+\/)*\w+\.\w+'\))(,\d+)?\)`)

	ParametrosLectorArchivo = util.Expresion{
		Regex:     exp,
		Categoria: "PARÁMETROS FUNCIÓN LECTURA DE ARCHIVOS",
	}

	util.VerificarError(err)
}
