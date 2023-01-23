package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "Konnichiwa, "
const world = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = world
	}
	var prefix string
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
