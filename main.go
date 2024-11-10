package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Token types
const (
	NUMBER   = "NUMBER"
	OPERATOR = "OPERATOR"
)

// Token represents a lexical token
type Token struct {
	Lexeme string
	Type   string
}

// isOperator checks if a character is a valid operator
func isOperator(ch rune) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}

	
}

// tokenize splits the input string into tokens
func tokenize(input string) []Token {
	var tokens []Token
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lexeme := scanner.Text()
		if _, err := strconv.Atoi(lexeme); err == nil {
			tokens = append(tokens, Token{Lexeme: lexeme, Type: NUMBER})
		} else if len(lexeme) == 1 && isOperator(rune(lexeme[0])) {
			tokens = append(tokens, Token{Lexeme: lexeme, Type: OPERATOR})
		} else {
			log.Fatalf("Unknown lexeme: %s", lexeme)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}

// readFile reads the content of a file
func readFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	var filePath string

	fmt.Println("What is the name of your file:?\nPlease add the extension (i.e: names.txt)")
	fmt.Scanln(&filePath)

	content := readFile(filePath)
	tokens := tokenize(content)

	fmt.Printf("%-10s %-10s\n", "Lexeme", "Token")
	fmt.Println(strings.Repeat("-", 20))
	for _, token := range tokens {
		fmt.Printf("%-10s %-10s\n", token.Lexeme, token.Type)
	}
}
