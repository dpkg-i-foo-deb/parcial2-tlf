package main

import (
	"analizador-lexico/analizador"
	"analizador-lexico/lector"
	"fmt"
)

func main() {

	lector.LeerArchivo()

	analizador.Datos = []byte(lector.Contenido)

	analizador.Inicializar()

	analizador.Leer()

	imprimir()

}

func imprimir() {

	fmt.Printf("Tipo de Token\tLexema\t\n")

	for _, t := range analizador.Tokens {
		fmt.Printf(t.Categoria + "\t" + t.Valor + "\n")
	}
}
