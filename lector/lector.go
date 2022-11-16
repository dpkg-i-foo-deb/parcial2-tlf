package lector

import (
	"fmt"
	"os"
)

var Contenido string

func LeerArchivo() {
	if len(os.Args) < 2 {
		fmt.Println("¿Qué se supone que voy a analizar? ¡Dame un archivo!")
		os.Exit(0)
	}

	datos, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Esto no parece un archivo que pueda leer...")
		os.Exit(1)
	}

	Contenido = string(datos)

}
