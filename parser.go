package main

import (
	"github.com/alecthomas/participle/v2"
)

type Expr struct {
	Op string `@(  "+" | "!" "=" | "/" "/" | "/" )`
}

var (
	parser = participle.MustBuild[Expr]()
)

func GetEBNF() string {
	return parser.String()
}
