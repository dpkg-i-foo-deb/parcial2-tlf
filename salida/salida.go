package salida

import (
	"analizador-lexico/analizador"
	"analizador-lexico/util"
	"bufio"
	"errors"
	"fmt"
	"os"
)

var archivo *os.File

func EscribirArchivo() {
	defer archivo.Close()

	if _, err := os.Stat("salida.txt"); errors.Is(err, os.ErrNotExist) {
		archivo, err = os.Create("salida.txt")

		util.VerificarError(err)
	}

	var err error

	archivo, err = os.Open("salida.txt")

	w := bufio.NewWriter(archivo)

	util.VerificarError(err)

	fmt.Fprintf(w, "Tipo de Token\tLexema\t\n")

	for _, t := range analizador.Tokens {
		fmt.Fprintf(w, t.Categoria+"\t"+t.Valor+"\n")
		fmt.Println(t.Categoria)
	}

	w.Flush()
}
