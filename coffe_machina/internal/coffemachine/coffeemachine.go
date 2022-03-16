package coffemachine

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/erikdzugkoev/coffee_task/models"
)

func Start(cm *models.CoffeeMachine) error {
	for {
		fmt.Println("Введите команду (buy, fill, take, remaining, exit):")
		reader := bufio.NewReader(os.Stdin)
		inp, err := reader.ReadString('\n')
		inp = strings.TrimSpace(inp)
		if err != nil {
			return err
		}

		switch inp { // подключаем свитч к обработке ввода пользователя
		case "buy":
			if err := Buy(cm); err != nil {
				fmt.Println("Не получилось приготовить кофе: " + err.Error())
			}
		case "fill":
			Fill(cm)
		case "take":
			if err := Take(cm); err != nil {
				fmt.Println("Не получилось вывести деньги: " + err.Error())
			}
		case "remaining":
			if err := Remaining(cm); err != nil {
				fmt.Println("nado popolnitm" + err.Error())
			}
		case "exit":
			os.Exit(0)
		default:

		}

	}
}

func Buy(cm *models.CoffeeMachine) error {

	fmt.Println("Выберите напиток: ")
	fmt.Println("Для выхода в главное меню нажмите *")
	for key, _ := range cm.CoffeeTypes {
		fmt.Println(key)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		inpCoffee, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		inpCoffee = strings.TrimSpace(inpCoffee)
		//inpCoffee1 = strconv.Atoi(inpCoffee)
		// Проверяем, существует ли такой кофе
		if inpCoffee == "*" {
			fmt.Println("Вы вернулись в главное меню")
			break
		}
		if err := cm.ValidateCoffeType(inpCoffee); err != nil {
			fmt.Println(err.Error())
			fmt.Println("Попробуй еще раз")
			continue
		}

		var coffee = cm.CoffeeTypes[inpCoffee]

		err = cm.Buy(coffee)
		if err != nil {
			return err
		}

		cm.AddCoffeeToStat(inpCoffee)

		for i := 0; i < coffee.WaitTime; i++ {
			fmt.Print(". ")
			time.Sleep(1 * time.Second)
		}

		fmt.Print("\nВаше кофе готов!\n")

		break
	}

	return nil

}

func Fill(cm *models.CoffeeMachine) error {

	milk := fillHelper("молока")
	water := fillHelper("воды")
	sugar := fillHelper("сахара")
	cups := fillHelper("чашек")
	coffeebeans := fillHelper("кофейных зерен")

	cm.Fill(milk, water, sugar, cups, coffeebeans)
	return nil
}

func fillHelper(text string) int {
	reader := bufio.NewReader(os.Stdin)
	var res int
exit:
	for {
		fmt.Println("Для возвращения в главное меню нажмите exit")
		fmt.Printf("Сколько добавить %s: ", text)
		inp, err := reader.ReadString('\n')
		if inp == "*" {
			inp = strings.TrimSpace(inp)
			break exit
		}
		if err != nil {
			fmt.Println("Произошла ошибка, попробуйте еще раз")
			continue
		}

		if inp != "" {
			res, err = strconv.Atoi(inp)
			if err != nil {
				fmt.Printf("Можно вводить только числа: %s\nПопробуйте еще раз\n", err.Error())
				continue
			}
		} else {
			res = 0
		}
		break
	}

	return res
}

func Take(cm *models.CoffeeMachine) error {
	money := takeHelper("Денег", cm.Money)

	if money <= cm.Money {
		cm.Take(money)
		fmt.Printf("В кассе осталось %d\n", cm.Money)
		return nil
	}
	return errors.New("вы ввели сумму, превыщающую остаток в кассе")

}

func takeHelper(textM string, texto int) int {
	reader := bufio.NewReader(os.Stdin)
	var ost int
	//var texto int
	for {
		fmt.Printf("Сколько хотите снять денег:\nВ кассе %s %dруб. \n", textM, texto)

		inp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Произашла ошибка, попробуйте ещё раз")
			continue
		}
		inp = strings.TrimSpace(inp)

		if inp != "" {
			ost, err = strconv.Atoi(inp)
			if err != nil {
				fmt.Printf("Можно вводить только числа: %s \n Попробуйте ещё раз\n", err.Error())
				continue
			}

		} else {
			ost = 0
		}
		break
	}
	return ost
}

func Remaining(cm *models.CoffeeMachine) error {
	milk := cm.Milk
	water := cm.Water
	sugar := cm.Sugar
	cups := cm.Cups
	coffeebeans := cm.CoffeeBeans
	money := cm.Money
	fmt.Printf("Сейчас в кофемашине: \nМолока %d мл \nВоды %d мл \nСахара %d грамм \nЧашек %d шт \nКофе %d грамм \nДенег %d руб\n\n", milk, water, sugar, cups, coffeebeans, money)

	fmt.Println("Всего проданно: ")
	for key, valaue := range cm.Stat {
		fmt.Printf("%s - %d\n", key, valaue)
	}

	return nil
}
