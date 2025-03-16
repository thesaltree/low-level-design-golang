package atm

import "errors"

type Account interface {
	GetBalance() float64
}

type AccountType string

const (
	Saving  AccountType = "SAVING"
	Current AccountType = "CURRENT"
)

func AccountFactory(accType AccountType) (Account, error) {
	switch accType {
	case Current:
		return &CurrentAccount{}, nil
	case Saving:
		return &SavingAccount{}, nil
	default:
		return nil, errors.New("Not a valid account type")
	}
}

type SavingAccount struct {
	acctNo  string
	balance float64
}

func (s *SavingAccount) GetBalance() float64 {
	return s.balance
}

type CurrentAccount struct {
	acctNo  string
	balance float64
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
