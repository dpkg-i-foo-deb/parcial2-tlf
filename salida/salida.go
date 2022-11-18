package salida

import (
	"analizador-lexico/analizador"
	"analizador-lexico/util"
	"errors"
	"fmt"
	"os"
)

var contenido []byte

func EscribirArchivo() {

	var s string

	if _, err := os.Stat("salida.txt"); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create("salida.txt")

		util.VerificarError(err)
	}

	s = fmt.Sprintf("Tipo de Token" + "\t" + "Lexema\t\n")

	contenido = append(contenido, []byte(s)...)

	for _, t := range analizador.Tokens {
		s = fmt.Sprintf(t.Categoria + "\t" + t.Valor + "\n")
		contenido = append(contenido, []byte(s)...)
	}

	os.WriteFile("salida.txt", contenido, 0644)
}
