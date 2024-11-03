package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	RAILROAD     = "/tmp/bin/railroad"
	EBNF_FILE    = "diagram/dsl.ebnf"
	DIAGRAM_FILE = "diagram/dsl.html"
)

func main() {
	genEbnf()
	genDiagram()
}

func genEbnf() {
	ebnf := GetEBNF()

	f, err := os.Create(EBNF_FILE)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(ebnf)
	if err != nil {
		log.Fatal(err)
	}

}

func genDiagram() {
	cmdRaw := fmt.Sprintf(`cat %s | %s -o %s`, EBNF_FILE, RAILROAD, DIAGRAM_FILE)
	out, err := exec.Command("bash", "-c", cmdRaw).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
