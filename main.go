package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"cryps/packages"
)

func main() {
	fmt.Println("1: Input text, get cypher\n2: Input cypher get text")

	var c int
	fmt.Scanln(&c)

	fmt.Println("Input file name(with .txt at the end): ")
	var d string
	fmt.Scanln(&d)

	content, err := os.ReadFile(d)
	if err != nil {
		log.Fatal(err)
	}
	var response []string

	if c == 1 {
		response = packages.Encrypt(content)
	} else {
		response = packages.Decrypt(content)
	}

	f, err := os.Create("response.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	key := strings.Join(response, "")
	_, err = f.WriteString(key)
	if err != nil {
		log.Fatal(err)
	}
}
