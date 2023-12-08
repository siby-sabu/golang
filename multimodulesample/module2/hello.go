package main

import (
	"fmt"
	"multimodulesample/module1"
)

func main() {
	msg := module1.Hello("Siby")
	fmt.Print(msg)
}
