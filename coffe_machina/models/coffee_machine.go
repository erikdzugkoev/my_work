package models

import (
	"errors"
)

type CoffeeMachine struct {
	Water       int
	Milk        int
	CoffeeBeans int
	Sugar       int
	Money       int
	Cups        int
	CoffeeTypes map[string]Coffee
	Stat        map[string]int
}

func NewCoffeeMachine(w, cb, m, c, mon, sug int) *CoffeeMachine {
	return &CoffeeMachine{
		Water:       w,
		CoffeeBeans: cb,
		Milk:        m,
		Cups:        c,
		Money:       mon,
		Sugar:       sug,
	}
}

func (cm *CoffeeMachine) InitCoffeTypes() {
	cappuch := Coffee{
		Milk:        111,
		Water:       0,
		CoffeeBeans: 10,
		Cost:        60,
		Sugar:       5,
		WaitTime:    5,
	}

	amerikano := Coffee{
		Milk:        0,
		Water:       100,
		CoffeeBeans: 15,
		Cost:        50,
		Sugar:       5,
		WaitTime:    3,
	}
	espresso := Coffee{
		Milk:        0,
		Water:       100,
		CoffeeBeans: 20,
		Cost:        40,
		Sugar:       5,
		WaitTime:    2,
	}

	coffeTypes := map[string]Coffee{
		"cappuch":   cappuch,
		"amerikano": amerikano,
		"espresso":  espresso,
	}
	cm.CoffeeTypes = coffeTypes

	cm.Stat = map[string]int{
		"cappuch":   0,
		"amerikano": 0,
		"espresso":  0,
	}
}

func (cm *CoffeeMachine) AddCoffeeToStat(inpCoffee string) {
	cm.Stat[inpCoffee] += 1
}

func (cm *CoffeeMachine) Buy(inp Coffee) error {
	if cm.Milk < inp.Milk {
		return errors.New("закончилось молоко")
	}
	if cm.Water < inp.Water {
		return errors.New("закончилась вода")
	}
	if cm.CoffeeBeans < inp.CoffeeBeans {
		return errors.New("закончилось кофе")
	}
	if cm.Cups < 0 {
		return errors.New("закончились стаканчики")
	}
	if cm.Sugar < inp.Sugar {
		return errors.New("закончился Сахар")
	}

	cm.Water -= inp.Water
	cm.Cups -= 1
	cm.Milk -= inp.Milk
	cm.Sugar -= inp.Sugar
	cm.Money += inp.Cost
	cm.CoffeeBeans -= inp.CoffeeBeans

	return nil
}

func (cm *CoffeeMachine) Fill(milk, water, sugar, coffeebeans, cups int) error {
	cm.Milk += milk
	cm.Water += water
	cm.Sugar += sugar
	cm.CoffeeBeans += coffeebeans
	cm.Cups += cups
	return nil
}

func (cm *CoffeeMachine) Take(money int) error {
	cm.Money -= money
	return nil
}

func (cn *CoffeeMachine) ValidateCoffeType(inp string) error {
	for coffeeType, _ := range cn.CoffeeTypes {
		if inp == coffeeType {
			return nil
		}
	}
	//a := strconv.Itoa(inp)
	return errors.New("Такого кофе нет: " + inp)
}
