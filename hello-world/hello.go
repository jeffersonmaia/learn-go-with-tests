package main

import "fmt"

//languages
const Spanish = "ES"
const French = "FR"

//prefix
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == Spanish {
		return spanishHelloPrefix + name
	}

	if language == French {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
