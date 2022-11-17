package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var Cadena util.Expresion

func InicializarCadenas() {

	//Creditos a Stiven Herrera
	exp, err := regexp.Compile(`("[^"]*")|('[^']*')`)

	Cadena = util.Expresion{
		Regex:     exp,
		Categoria: "CADENA",
	}

	util.VerificarError(err)
}
