package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var OperadorLogico util.Expresion
var OperadorAsignacion util.Expresion
var OperadorRelacional util.Expresion
var OperadorAritmetico util.Expresion
var OperadorTerminal util.Expresion
var UltimoIndice util.Expresion

func InicializarOperadores() {

	exp, err := regexp.Compile(`[|&~]`)

	util.VerificarError(err)

	OperadorLogico = util.Expresion{
		Regex:     exp,
		Categoria: "OPERADOR_LOGICO",
	}

	exp, err = regexp.Compile(`[=]`)

	OperadorAsignacion = util.Expresion{
		Regex:     exp,
		Categoria: "OPERADOR_ASIGNACION",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`(>=|<=|<|>|==|~=)`)

	OperadorRelacional = util.Expresion{
		Regex:     exp,
		Categoria: "OPERADOR_RELACIONAL",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`[+\-*\\\/\^]`)

	OperadorAritmetico = util.Expresion{
		Regex:     exp,
		Categoria: "OPERADOR_ARITMETICO",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`[;]`)

	OperadorTerminal = util.Expresion{
		Regex:     exp,
		Categoria: "OPERADOR_TERMINAL",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`[$]`)

	UltimoIndice = util.Expresion{
		Regex:     exp,
		Categoria: "ULTIMO_INDICE",
	}

	util.VerificarError(err)
}
