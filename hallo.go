package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var p1 = fmt.Println()

func main() {
	//p1("Hallo Go")
	p1("Wie heißt du?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		p1("Hallo", name)
	} else {
		log.Fatal(err)
	}
}
