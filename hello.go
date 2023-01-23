package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name != "" {
		return englishHelloPrefix + name
	} else {
		return englishHelloPrefix + "World"
	}
}

func main() {
	fmt.Println(Hello("world"))
}