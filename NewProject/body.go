package body

import (
	"bufio"
	"fmt"
	"os"
)

func start(p string) error {
	for {
		
		fmt.Println("Привет, ты пришел с новой прочитанной книгой ?(Да/Нет")
		reader := bufio.NewReader(os.Stdin)
		Inp, err := reader.ReadString("\n")
		if err != nil {
			return err
		}
		switch Inp {
		case "Да":
			if err := Yes(); err != nil {
				fmt.Print("Что-то пошло не так")
			}
		сase "Нет":
			os.Exit(0)	
		}
	}
}
func Yes(yes string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Очень классно, давай сохраним инфу о ней, Введи Название: ")
		inp,_ := reader.ReadString("\n")

	}
}
