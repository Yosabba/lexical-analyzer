package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	IDENTIFIER  = "identifier"
	INT_LITERAL = "int_literal"
	OPERATOR    = "operator"
	EQUALS      = "equal_sign"
	DELIMITER   = "delimiter"
)

type Token struct {
	Lexeme string
	Type   string
}

func isCharOperator(character rune) bool {
	switch character {
	case '+', '-', '*', '/', '=':
		return true
	default:
		return false
	}
}

func isCharDelimiter(character rune) bool {
	switch character {
	case '(', ')', '{', '}', '[', ']', ';', ',', '_':
		return true
	default:
		return false
	}
}

func isNumber(word string) bool {
	for _, char := range word {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isIdentifier(word string) bool {
	if len(word) == 0 {
		return false
	}
	for i, char := range word {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char != '_') && (i == 0 || (char < '0' || char > '9')) {
			return false
		}
	}
	return true
}

func tokenizeFile(input string) []Token {
	var tokens []Token
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if isNumber(word) {
			tokens = append(tokens, Token{Lexeme: word, Type: INT_LITERAL})
		} else if len(word) == 1 {
			char := rune(word[0])
			switch {
			case isCharOperator(char):
				if word == "=" {
					tokens = append(tokens, Token{Lexeme: word, Type: EQUALS})
				} else {
					tokens = append(tokens, Token{Lexeme: word, Type: OPERATOR})
				}
			case isCharDelimiter(char):
				tokens = append(tokens, Token{Lexeme: word, Type: DELIMITER})
			}
		} else if isIdentifier(word) {
			tokens = append(tokens, Token{Lexeme: word, Type: IDENTIFIER})
		}
	}
	return tokens
}

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
