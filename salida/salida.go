package salida

import (
	"analizador-lexico/analizador"
	"analizador-lexico/util"
	"errors"
	"fmt"
	"os"
	"strings"
)

var contenido []byte

func EscribirArchivo() {
	linea := obtenerCabecera()
	contenido = append(contenido, []byte(linea)...)

	if _, err := os.Stat("salida.txt"); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create("salida.txt")

		util.VerificarError(err)
	}

	for _, t := range analizador.Tokens {
		linea = fmt.Sprintf("| %-40s | %-50s |\n", t.Categoria, t.Valor)
		contenido = append(contenido, []byte(linea)...)
	}

	err := os.WriteFile("salida.txt", contenido, 0644)

	util.VerificarError(err)

}

func obtenerCabecera() string {
	separador := "|" + strings.Repeat("-", 42) + "|" + strings.Repeat("-", 52) + "|\n"

	cabecera := separador
	cabecera += fmt.Sprintf("| %-40s |", strings.Repeat(" ", 12)+"Tipo de Token")
	cabecera += fmt.Sprintf(" %-50s |\n", strings.Repeat(" ", 20)+"Lexema")
	cabecera += separador

	return cabecera
}
