package main

import (
	"analizador-lexico/analizador"
	"fmt"
)

func main() {

	analizador.Texto = `variable1 = 2;
				variable2 = 3;
				variable4=3+4;
				//Esto es un comentario
				variable8="Cadena"
				if (4>5) then
				
					else 

						if (variable2<variable5) then
							print('holis.txt',x1)
						end

				end`

	analizador.Inicializar()

	analizador.Datos = []byte(analizador.Texto)

	for ok := true; ok; ok = len(analizador.Datos) > 0 {
		analizador.Analizar()
	}

	imprimir()

}

func imprimir() {

	fmt.Printf("Tipo de Token\tLexema\t\n")

	for _, t := range analizador.Tokens {
		fmt.Printf(t.Categoria + "\t" + t.Valor + "\n")
	}
}
