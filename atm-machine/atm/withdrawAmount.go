package atm

import (
	"errors"
	"fmt"
)

type DispenserAmount struct {
	atm *ATM
	ATMAbstract
}

func (w *DispenserAmount) DispenserAmount() error {
	var amount float64

	fmt.Println("\n Enter amount to withdraw from atm:")
	fmt.Scanf("%f", &amount)
	// w.atm.amount = amount
	out := NewWithDrawPipeline()

	out.ProcessAmount(w.atm, amount) // this process the amount

	if w.atm.WithdrawAs.Left > 0 {
		return errors.New("This amount is not divisible of 100")
	} else {
		if w.atm.countOFNotes["500"] >= w.atm.WithdrawAs.FiveHundread {
			w.atm.countOFNotes["500"] = w.atm.countOFNotes["500"] - w.atm.WithdrawAs.FiveHundread
		} else {
			return errors.New("Bank don't have enough fund 500")
		}

		if w.atm.countOFNotes["100"] >= w.atm.WithdrawAs.Hundread {
			w.atm.countOFNotes["100"] = w.atm.countOFNotes["100"] - w.atm.WithdrawAs.Hundread
		} else {
			return errors.New("Bank don't have enough fund 100")
		}
	}

	w.atm.ResetAtm()
	w.atm.SetState(w.atm.insertCard)

	return nil
}

func (w *DispenserAmount) StateName() string {
	return "DispenserAmount"
}
