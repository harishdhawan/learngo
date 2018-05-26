package main

import "fmt"

type Printer func(string)

type Salutation struct {
	name    string
	message string
}

func Print1(s string) {
	fmt.Print(s)
}

func Print2(s string) {
	fmt.Println(s)
}

func CreatePrinter(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func GoPrint(s Salutation, printer Printer) {
	printer(s.name)
	printer(s.message)
}

func CreateMessage(salutation Salutation) string {
	return salutation.name + " " + salutation.message
}

func main() {
	salutation := Salutation{"Harish", "Hello"}
	GoPrint(salutation, CreatePrinter("! "))
}
