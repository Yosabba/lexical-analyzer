package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Token struct {
	Lexeme string
	Type   string
}

func isCharOperator(character rune) bool {
	switch character {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

// tokenize splits the input string into tokens
func tokenizeFile(input string) []Token {
	var tokens []Token

	NUMBER := "NUMBER"
	OPERATOR := "OPERATOR"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lexeme := scanner.Text()
		if _, err := strconv.Atoi(lexeme); err == nil {
			tokens = append(tokens, Token{Lexeme: lexeme, Type: NUMBER})
		} else if len(lexeme) == 1 && isCharOperator(rune(lexeme[0])) {
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
