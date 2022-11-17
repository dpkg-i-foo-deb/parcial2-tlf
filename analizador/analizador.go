package analizador

import (
	"analizador-lexico/expresiones"
	"analizador-lexico/util"
	"strings"
)

var Patrones []util.Expresion
var Datos []byte
var Texto string
var Tokens []util.Token

func Analizar() {

	var identificado bool
	//Pasamos el Texto por cada expresión regular
	for contadorPatrones := 0; contadorPatrones < len(Patrones); contadorPatrones++ {

		if Patrones[contadorPatrones].Regex.Match(Datos) {

			indices := Patrones[contadorPatrones].Regex.FindIndex(Datos)

			//Es vital que no sólo coincida el patrón, también debe estar en la primera posición
			if indices[0] != 0 {
				continue
			}
			//Removemos la ocurrencia de los Datos en forma de Texto
			encontrado := Patrones[contadorPatrones].Regex.Find([]byte(Texto))
			Texto = strings.Replace(Texto, string(encontrado), "", 1)

			//Categorizamos el token encontrado
			t := util.Token{
				Valor:     string(encontrado),
				Categoria: Patrones[contadorPatrones].Categoria,
			}

			//Si es un salto de línea, es ignorado
			if Patrones[contadorPatrones].Categoria != "SALTO_LINEA" {
				Tokens = append(Tokens, t)
			}

			//Ahora cambiamos los Datos al nuevo Texto sin la ocurrencia
			Datos = []byte(Texto)

			identificado = true

		}

	}

	if !identificado {
		//Si no coincide con nada, entonces no es un token valido
		t := util.Token{
			Valor:     string(Datos[0]),
			Categoria: "NO VALIDO",
		}

		Tokens = append(Tokens, t)
		//Elimino la posicion actual de los Datos
		Datos = append(Datos[:0], Datos[1:]...)
		Texto = string(Datos)
		//Ahora cambiamos los Datos al nuevo Texto sin la ocurrencia
		Datos = []byte(Texto)
	}
}

func Inicializar() {
	compilarExpresiones()
	Texto = IgnorarCampos(Texto)
}

func compilarExpresiones() {
	expresiones.InicializarOperadores()
	expresiones.InicializarSeparadorSentencias()
	expresiones.InicializarNumeros()
	expresiones.InicializarNombres()
	expresiones.InicializarComentarios()
	expresiones.InicializarCadenas()
	expresiones.InicializarBasura()
	expresiones.InicializarCondicionales()
	expresiones.InicializarImpresion()
	expresiones.InicializarCiclos()

	//Agregamos todas las expresiones regulares a un arreglo con un orden
	Patrones = append(Patrones, expresiones.SaltoLinea)
	Patrones = append(Patrones, expresiones.CicloFor)
	Patrones = append(Patrones, expresiones.OperacionImprimir)
	Patrones = append(Patrones, expresiones.CondicionalIf)
	Patrones = append(Patrones, expresiones.CondicionalElse)
	Patrones = append(Patrones, expresiones.CondicionalElseIf)
	Patrones = append(Patrones, expresiones.CicloWhile)
	Patrones = append(Patrones, expresiones.PalabraReservadaEnd)
	Patrones = append(Patrones, expresiones.PalabraReservadaThen)
	Patrones = append(Patrones, expresiones.Comentarios)
	Patrones = append(Patrones, expresiones.ComentariosComplejos)
	Patrones = append(Patrones, expresiones.OperadorAsignacion)
	Patrones = append(Patrones, expresiones.Entero)
	Patrones = append(Patrones, expresiones.Decimal)
	Patrones = append(Patrones, expresiones.SeparadorSentencias)
	Patrones = append(Patrones, expresiones.OperadorLogico)
	Patrones = append(Patrones, expresiones.OperadorRelacional)
	Patrones = append(Patrones, expresiones.OperadorAritmetico)
	Patrones = append(Patrones, expresiones.OperadorTerminal)
	Patrones = append(Patrones, expresiones.UltimoIndice)
	Patrones = append(Patrones, expresiones.Cadena)
	Patrones = append(Patrones, expresiones.Parametros_Condicional)
	Patrones = append(Patrones, expresiones.ParametroCicloFor)
	Patrones = append(Patrones, expresiones.ParametrosImprimir)
	Patrones = append(Patrones, expresiones.NombreVariables)
}

func IgnorarCampos(s string) string {

	//Ignoramos los espacios y tabulaciones
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	//s = strings.Replace(s, "\n", "", -1)

	return s

}
