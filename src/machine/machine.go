package machine

import (
	"errors"
	"fmt"
	"time"
)

type MachineState int

const (
	Idle MachineState = iota
	Working
)

type BeverageType int

const (
	Coffee BeverageType = iota + 1
	SoftDrink
)

type VendingMachine struct {
	Id                  string
	state               MachineState
	StatusChangedToIdle chan bool
	order               chan bool
	coffeeQuantity      int
	softDrinkQuantity   int
	cashBalance         int
}

func NewVendingMachine(id string, coffeeQuantity, softDrinkQuantity int) *VendingMachine {
	m := &VendingMachine{
		Id:                  id,
		state:               Idle,
		StatusChangedToIdle: make(chan bool),
		order:               make(chan bool),
		coffeeQuantity:      coffeeQuantity,
		softDrinkQuantity:   softDrinkQuantity,
	}
	//start the machine
	go m.run()
	return m
}

func (m *VendingMachine) PlaceOrder() {
	if m.state == Working {
		fmt.Printf("Machine #%s is not available now, use another or try again later \n", m.Id)
		return
	}
	m.order <- true
}

func (m *VendingMachine) run() {
	for {
		select {
		case <-m.order:
			m.state = Working
			selection, err := m.selectBeverage()
			if err != nil {
				fmt.Println(err.Error())
				m.state = Idle
				m.StatusChangedToIdle <- true
				continue
			}
			err = m.insertCoin(selection)
			if err != nil {
				fmt.Println(err.Error())
			}
			m.state = Idle
			m.StatusChangedToIdle <- true
		}
	}
}

func (m *VendingMachine) insertCoin(selection BeverageType) error {
	fmt.Println("Please insert a coin:[Just type 1 and press enter]")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil || input != 1 {
		return errors.New("error in reading or invalid coin")
	}
	m.cashBalance++
	if selection == Coffee {
		m.coffeeQuantity--
		fmt.Printf("Your coffee will be ready soon. Please wait\n")
		time.Sleep(4 * time.Second)
		fmt.Printf("Here you are: COFFEE\n\n")
	} else if selection == SoftDrink {
		fmt.Printf("Your drink will be ready soon. Please wait\n")
		m.softDrinkQuantity--
		time.Sleep(1 * time.Second)
		fmt.Printf("Here you are: SOFTDRINK\n\n")
	}
	return nil
}

func (m *VendingMachine) selectBeverage() (BeverageType, error) {
	fmt.Printf("Machine #%s is at your service, please select a beverage:\n 1.Coffee	2.Soft Drink\n", m.Id)
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil || (input != 1 && input != 2) {
		return 0, errors.New("error in reading or invalid option")
	}
	beverage := BeverageType(input)
	if beverage == Coffee {
		if 1 <= m.coffeeQuantity {
			return Coffee, nil
		} else {
			return 0, errors.New("coffee is not available now")
		}
	} else {
		if 1 <= m.softDrinkQuantity {
			return SoftDrink, nil
		} else {
			return 0, errors.New("soft drink is not available now")
		}
	}
}
