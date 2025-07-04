package main

import (
	"fmt"
	"log"
	// "bufio"
	// "os"
	// "net/http"
	"flag"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	flag.Parse()
	fmt.Printf("---TCP SCANNER---\n")
}