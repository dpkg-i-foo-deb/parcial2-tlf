package salida

import (
	"analizador/lector"
	"errors"
	"os"
)

func EscribirArchivo() {
	if _, err := os.Stat("salida.txt"); errors.Is(err, os.ErrNotExist) {
		os.Create("salida.txt")
	}

	os.WriteFile("salida.txt", []byte(lector.Contenido), 0644)
}
