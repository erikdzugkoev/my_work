package main

import (
	"fmt"
	"os"

	"github.com/erikdzugkoev/coffee_task/internal/coffemachine"
	"github.com/erikdzugkoev/coffee_task/models"
)
func main() {
	cm := models.NewCoffeeMachine(400, 120, 540, 9, 550)
	cm.InitCoffeTypes()

	if err := coffemachine.Start(cm); err != nil { // Для чего здесь эрор
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
