package atm

import (
	"errors"
	"fmt"
)

type InsertPin struct {
	atm *ATM
	ATMAbstract
}

func (s *InsertPin) InsertPin() error {
	// logic to valid pin, we can call api client inside it, i will return sucess and error depend on input
	fmt.Println("Enter 4 letter pin for card:\n")

	var pin string

	fmt.Scanf("%s", &pin)

	if pin == "1111" {
		return nil
	} else {
		return errors.New("Pin is not valid")
	}
}

func (d *InsertPin) StateName() string {
	return "InsetPin"
}
