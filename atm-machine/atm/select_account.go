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

	s.atm.account, err = AccountFactory(AccountType(accountType))

	s.atm.SetState(s.atm.dispenserAmount)

	return err

}

func (s *SelectAccount) StateName() string {
	return "SelectAccount"
}
