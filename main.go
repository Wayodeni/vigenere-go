package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var alphabet string
	var key string
	var message string

	flag.StringVar(&alphabet, "alphabet", "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ", "all alphabet symbols")
	flag.StringVar(&key, "key", "ЛЕВ", "keyphrase to encrypt data")
	flag.StringVar(&message, "message", "В КНИГАХ МНОГО ТАЙН И РАЗГАДОК", "message to encrypt")

	flag.Parse()

	encryptedMessage, err := EncryptVigenere(alphabet, message, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(encryptedMessage)
}

func EncryptVigenere(alphabet, message, key string) (string, error) {
	keyRunes := []rune(strings.ToUpper(key))
	alphabetRunes := []rune(strings.ToUpper(alphabet))
	messageRunes := []rune(strings.ToUpper(message))
	var encryptedMessage string

	for messageRuneIndex := 0; messageRuneIndex < len(messageRunes); messageRuneIndex++ {
		messageRune := messageRunes[messageRuneIndex]

		keyRune := keyRunes[messageRuneIndex%len(keyRunes)]

		symbolShift, ok := IndexOf[rune](keyRune, alphabetRunes)
		if !ok {
			return "", fmt.Errorf("unknown symbol '%s' in key", string(keyRune))
		}

		indexOfMessageSymbolInAlphabet, ok := IndexOf[rune](messageRune, alphabetRunes)
		if !ok {
			return "", fmt.Errorf("unknown symbol '%s' in message", string(messageRune))
		}

		encryptedSymbolIndex := indexOfMessageSymbolInAlphabet + symbolShift
		if encryptedSymbolIndex > len(alphabetRunes) {
			encryptedSymbolIndex = encryptedSymbolIndex % len(alphabetRunes)
		}

		encryptedSymbol := string(alphabetRunes[encryptedSymbolIndex])

		encryptedMessage += encryptedSymbol
	}

	return encryptedMessage, nil
}

func IndexOf[T comparable](element T, slice []T) (index int, ok bool) {
	for index, sliceElement := range slice {
		if element == sliceElement {
			return index, true
		}
	}

	return -1, false
}
