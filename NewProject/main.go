package main

import (
	"fmt"
	"os"
)

func main() {
	start("asas")
	if err := start(); err != nil { // Для чего здесь эрор
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
