package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-a/pkgmoda"
)

func main() {
	fmt.Println("Hello World")
	pkgmoda.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-a")
}
