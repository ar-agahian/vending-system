package main

import (
	"fmt"
	"github.com/ar-agahian/vending-system/machine"
)

func main() {
	machines := make(map[string]*machine.VendingMachine)
	for i := 1; i < 4; i++ {
		m := machine.NewVendingMachine(fmt.Sprint(i), 2, 2)
		machines[m.Id] = m
	}
	var input string
	// var pause chan bool
	for {
		fmt.Println("Select vending machine number (1-3) or press q to quit")
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
		vm, ok := machines[input]
		if !ok {
			fmt.Println("Machine number you entered is invalid")
			continue
		}
		vm.PlaceOrder()
		<-vm.StatusChangedToIdle
	}
}
