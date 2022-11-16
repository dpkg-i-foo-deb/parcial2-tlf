package main

import (
	"analizador/lector"
	"analizador/lexer_framework"
	"analizador/salida"
	"analizador/tokens"

	"fmt"
	"log"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

// NO TOCAR
func main() {

	lexer_framework.InitLexer()

	compilarTokens()

	compilarExpresiones()

	lexer_framework.Compilar()

	lector.LeerArchivo()

	analizar()

	salida.EscribirArchivo()

}

func compilarExpresiones() {
	tokens.AgregarComentarios()
	tokens.AgregarOperadores()
	tokens.AgregarSeparadorSentencias()
}

func compilarTokens() {
	var elementos []string

	elementos = append(elementos, tokens.Operadores...)
	elementos = append(elementos, tokens.Comentarios...)
	elementos = append(elementos, tokens.SeparadorSentencias...)

	lexer_framework.TokenIds = make(map[string]int)
	for i, tok := range elementos {
		lexer_framework.TokenIds[tok] = i
	}

	//Esto se hace de último
	lexer_framework.Tokens = elementos
}

func analizar() {
	s, err := lexer_framework.Lexer.Scanner([]byte(lector.Contenido))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tipo   		 | Lexema   		  |Posición		")
	fmt.Println("--------+------------+------------")
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if ui, is := err.(*machines.UnconsumedInput); is {
			// to skip bad token do:
			s.TC = ui.FailTC
			log.Fatal(err)
			//fmt.Print("No reconocido")
			break

		} else if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lexmachine.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			lexer_framework.Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
