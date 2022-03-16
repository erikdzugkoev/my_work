package main

import (
	"fmt"
	"os"

	"github.com/erikdzugkoev/coffee_task/internal/coffemachine"
	"github.com/erikdzugkoev/coffee_task/models"
)

func main() {
	cm := models.NewCoffeeMachine(400, 120, 540, 0, 550, 100)
	cm.InitCoffeTypes()

	if err := coffemachine.Start(cm); err != nil { 
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
