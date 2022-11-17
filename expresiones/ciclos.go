package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var CicloFor util.Expresion
var CicloWhile util.Expresion
var ParametroCicloFor util.Expresion
var ParametroCicloWhile util.Expresion

func InicializarCiclos() {
	// Sentencia for
	exp, err := regexp.Compile(`for`)

	CicloFor = util.Expresion{
		Regex: exp,
		Categoria: "CICLO FOR",
	}

	util.VerificarError(err)

	// Parámetros del for
	exp, err = regexp.Compile(`\(?([a-z]|[#%!$?_][^A-Z])([a-zA-Z]*|[0-9]*|[#%!$?_]*)*=.+(:.+)?:.+\)?`)

	ParametroCicloFor = util.Expresion{
		Regex: exp,
		Categoria: "Parámetro de ciclo",
	}

	util.VerificarError(err)

	// Sentencia while
	exp, err = regexp.Compile(`while`)

	CicloWhile = util.Expresion{
		Regex: exp,
		Categoria: "CICLO WHILE",
	}

	util.VerificarError(err)
}