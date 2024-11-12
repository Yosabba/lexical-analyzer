package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readUserFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	var fileName string

	fmt.Println("What is the name of your file:?\nPlease add the extension (i.e: names.txt)")
	fmt.Scanln(&fileName)

	contentOfFile := readUserFile(fileName)
	tokens := tokenizeFile(contentOfFile)

	//lexeme & token needs to be be in a table so format it to be in a table with a 10 char min
	fmt.Printf("%-10s %-10s\n", "Lexeme", "Token")
	fmt.Println(strings.Repeat("-", 20))

	for _, token := range tokens {
		fmt.Printf("%-10s %-10s\n", token.Lexeme, token.Type)
	}

}
