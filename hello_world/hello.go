package main

import "fmt"

const (
	spanish  = "Spanish"
	french   = "French"
	romanian = "Romanian"

	spanishHelloPrefix  = "Hola, "
	englishHelloPrefix  = "Hello, "
	frenchHelloPrefix   = "Bonjour, "
	romanianHelloPrefix = "Salut, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
