package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var Entero util.Expresion
var Decimal util.Expresion

func InicializarNumeros() {
	exp, err := regexp.Compile(`([0-9]|-[0-9])[0-9]*`)

	Entero = util.Expresion{
		Regex:     exp,
		Categoria: "ENTERO",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`([0-9]|-[0-9])[0-9]*\.[0-9][0-9]*`)

	Decimal = util.Expresion{
		Regex:     exp,
		Categoria: "DECIMAL",
	}

	util.VerificarError(err)
}
