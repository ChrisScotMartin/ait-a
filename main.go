package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-a/pkgmoda"
	"github.com/ChrisScotMartin/ait-b/pkgmodb"
	"github.com/ChrisScotMartin/ait-c/pkgmodc"
)

// ait-a imports pkg/mods from ait-b and ait-c
func main() {
	fmt.Println("Hello World")
	pkgmoda.Hello()
	pkgmodb.Hello()
	pkgmodc.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-a")
}
