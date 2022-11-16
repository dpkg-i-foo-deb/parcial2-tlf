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

	tokens.InicializaPalabrasReservadas()

	s, err := lexer_framework.Lexer.Scanner([]byte(`$`))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Type    | Lexeme     | Position")
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

	lector.LeerArchivo()

	salida.EscribirArchivo()

}
