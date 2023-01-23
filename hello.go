package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const world = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = world
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
