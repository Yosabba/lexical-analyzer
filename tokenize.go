package main

import (
	"bufio"

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
