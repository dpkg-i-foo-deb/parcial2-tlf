package salida

import (
	"analizador-lexico/analizador"
	"analizador-lexico/util"
	"errors"
	"os"
)

var contenido []byte

func EscribirArchivo() {

	contenido = append(contenido, []byte("Tipo de Token"+"\t"+"Lexema\t\n")...)

	if _, err := os.Stat("salida.txt"); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create("salida.txt")

		util.VerificarError(err)
	}

	for _, t := range analizador.Tokens {
		contenido = append(contenido, []byte(t.Categoria+"\t"+t.Valor+"\n")...)
	}

	err := os.WriteFile("salida.txt", contenido, 0644)

	util.VerificarError(err)

}
