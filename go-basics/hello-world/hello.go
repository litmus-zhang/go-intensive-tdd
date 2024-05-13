package main

import (
	"fmt"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}
	return
}
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func main() {
	fmt.Println(Hello("", ""))
}
