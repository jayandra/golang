package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

// ASK: set the default language to be empty string such that we can have only one argument
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix
}

func iterateHello(n int, name string, language string) string {
	retval := ""
	for i := 0; i < n; i++ {
		retval += Hello(name, language) + "\n"
	}
	return retval
}

func main() {
	// var name string
	// var language string
	// var times int

	// fmt.Println("Enter the name for greeting:")
	// fmt.Scanln(&name)
	// fmt.Println("Enter the language [Spanish, French, English(default)]")
	// fmt.Scanln(&language)
	// fmt.Println("Enter the times for repetition:")
	// fmt.Scanln(&times)
	// fmt.Println(iterateHello(times, name, language))

	// fmt.Println("Enter how many numbers you want to sum:")
	// fmt.Println("Enter the numbers:")

	numbers := []float64{1.0}
	// fmt.Scanln(&numbers)
	for i := 0; i < 2; i++ {
		var a float64
		fmt.Scanln(&a)
		numbers = append(numbers, a)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%.2f", numbers[i])
	}

	fmt.Printf("The sum of numbers is: %f", Sum(numbers))
}
