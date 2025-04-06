package atm

import (
	"fmt"
)

type SelectAccount struct {
	atm *ATM
	ATMAbstract
}

func (s *SelectAccount) SelectAccount() error {

	var accountType string
	var err error

	fmt.Println("Select account type SAVING OR CURRENT:\n")
	fmt.Scanf("%s", &accountType)

	if s.atm.account, err = AccountFactory(AccountType(accountType)); err != nil {
		return err
	}

	s.atm.SetState(s.atm.dispenserAmount)

	return nil

}

func (s *SelectAccount) StateName() string {
	return "SelectAccount"
}
