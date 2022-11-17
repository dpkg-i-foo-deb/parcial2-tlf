package expresiones

import (
	"analizador-lexico/util"
	"regexp"
)

var Comentarios util.Expresion
var ComentariosComplejos util.Expresion

func InicializarComentarios() {
	exp, err := regexp.Compile(`//[^\n]*\n?`)

	Comentarios = util.Expresion{
		Regex:     exp,
		Categoria: "COMENTARIO",
	}

	util.VerificarError(err)

	exp, err = regexp.Compile(`\/\*([^*]|\r|\n|(\*+([^*\/]|\r|\n)))*\*\/`)

	ComentariosComplejos = util.Expresion{
		Regex:     exp,
		Categoria: "COMENTARIO",
	}

	util.VerificarError(err)

}
