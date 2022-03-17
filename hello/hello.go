package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Chris", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	switch language {
	case french:
		return frenchHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}
