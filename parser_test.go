package main

import (
	"log"
	"testing"

	require "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
)

func TestOperationParser(t *testing.T) {
	type testCase struct {
		src      string
		expected Expr
	}

	for _, c := range []testCase{
		{`+`, Expr{Op: "+"}},   // ok
		{`!=`, Expr{Op: "!="}}, // ok
		{`&&`, Expr{Op: "&&"}}, // ok
		{`/`, Expr{Op: "/"}},   // ok
		{`//`, Expr{Op: "//"}}, // fails
	} {
		// raw parse test: without precedence
		// that parser handles all expected cases
		log.Printf("Testing: %v", c.src)
		actual, err := parser.ParseString("", c.src)
		repr.Println(*actual)
		require.NoError(t, err)
		require.Equal(t, c.expected, *actual, c.src)
	}
}
