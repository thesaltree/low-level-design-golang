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

type BaseAccount struct {
	acctNo  string
	balance float64
}

func (b *BaseAccount) GetBalance() float64 {
	return b.balance
}

type SavingAccount struct {
	BaseAccount
}

type CurrentAccount struct {
	BaseAccount
}
