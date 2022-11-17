package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var CondicionalIf util.Expresion
var CondicionalElse util.Expresion
var CondicionalElseIf util.Expresion
var PalabraReservadaThen util.Expresion
var PalabraReservadaEnd util.Expresion
var Parametros_Condicional util.Expresion

func InicializarCondicionales() {
	exp, err := regexp.Compile(`if`)

	CondicionalIf = util.Expresion{
		Regex:     exp,
		Categoria: "CONDICIONAL_IF",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`else`)

	CondicionalElse = util.Expresion{
		Regex:     exp,
		Categoria: "CONDICIONAL_ELSE",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`elseif`)

	CondicionalElseIf = util.Expresion{
		Regex:     exp,
		Categoria: "CONDICIONAL_ELSEIF",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`then`)

	PalabraReservadaThen = util.Expresion{
		Regex:     exp,
		Categoria: "PALABRA_RESERVADA_THEN",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`end`)

	PalabraReservadaEnd = util.Expresion{
		Regex:     exp,
		Categoria: "PALABRA_RESERVADA_END",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`\(?([0-9]|[a-z-A-Z])+(<=|==|>=|>|<|&&)([0-9]|[a-z-A-Z])+\)?`)

	Parametros_Condicional = util.Expresion{
		Regex:     exp,
		Categoria: "PARAMETROS_CONDICIONAL",
	}

	util.VerificarError(err)

}
