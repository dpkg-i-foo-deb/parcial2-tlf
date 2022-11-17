package main

import (
	"analizador-lexico/analizador"
	"analizador-lexico/lector"
	"analizador-lexico/salida"
)

func main() {

	lector.LeerArchivo()

	analizador.Datos = []byte(lector.Contenido)

	analizador.Inicializar()

	analizador.Leer()

	salida.EscribirArchivo()

}
