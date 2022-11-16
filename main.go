package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-a/pkgmoda"
	"github.com/ChrisScotMartin/ait-b"
	"github.com/ChrisScotMartin/ait-c/renamedpkgmodc"
)

// ait-a imports pkg/mods from ait-b and ait-c
func main() {
	fmt.Println("Hello World")
	pkgmoda.Hello()
	renamedpkgmodc.Hello()
	libb.SomeLibraryFunc()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-a")
}
