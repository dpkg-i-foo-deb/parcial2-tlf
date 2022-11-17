package util

import "regexp"

type Expresion struct {
	Regex     *regexp.Regexp
	Categoria string
}
