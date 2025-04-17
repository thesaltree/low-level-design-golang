package main

import (
	"atm-machine/atm"
	"fmt"
)

// flow of ATM
// insertCard state -> readCard state -> insertPin state -> selectAccount state -> withdrawAmount state
func main() {
	atm := atm.NewATM()

	for {

		var runProcess string

		fmt.Printf("\nATM state:%s\n", atm.StateName())
		fmt.Println("Input do you want to proceed, press n to exit")
		fmt.Scanf("%s", &runProcess)

		fmt.Println(runProcess)
		if runProcess == "n" {
			return
		}
		//
		fmt.Printf("\nATM state:%s\n", atm.StateName())
		// insert card
		atm.Execute(atm.InsertCard)

		fmt.Printf("\nATM state:%s\n", atm.StateName())

		atm.Execute(atm.GetCardDetail)

		fmt.Printf("\nATM state:%s\n", atm.StateName())

		atm.Execute(atm.InsertPin)

		fmt.Printf("\nATM state:%s\n", atm.StateName())
		// insert card
		atm.Execute(atm.SelectAccount)

		fmt.Printf("\nATM state:%s", atm.StateName())

		atm.PrintMoney()
		// insert card
		atm.Execute(atm.DispenserAmount)

		atm.PrintMoney()

		fmt.Printf("\nATM state:%s", atm.StateName())

	}

}
