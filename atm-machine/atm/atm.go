package atm

import (
	"atm-machine/model"
	"fmt"
	"sync"
)

type WithdrawNote struct {
	FiveHundread int
	Hundread     int
	Left         int
}

type ATM struct {
	countOFNotes    map[string]int
	card            model.Card
	account         Account
	uiOption        []string
	WithdrawAs      *WithdrawNote
	insertCard      ATMState
	readCard        ATMState
	selectAccount   ATMState
	dispenserAmount ATMState
	currentState    ATMState
	mu              sync.RWMutex
}

func (a *ATM) SetState(s ATMState) {
	a.currentState = s
}

func (a *ATM) ResetAtm() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.WithdrawAs = &WithdrawNote{}
	a.card = model.Card{}
	a.account = nil
}

func (a *ATM) PrintMoney() {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Printf("\n500 note present:%d, 100 not present:%d", a.countOFNotes["500"], a.countOFNotes["100"])
}

func (a *ATM) StateName() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.currentState.StateName()
}

func (a *ATM) InsertCard() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.currentState.InsertCard()
}

func (a *ATM) GetCardDetail() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.currentState.GetCardDetail()
}

func (a *ATM) DispenserAmount() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.currentState.DispenserAmount()
}

func (a *ATM) SelectAccount() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.currentState.SelectAccount()
}

func (a *ATM) Execute(operation func() error) {
	err := operation()
	if err != nil {
		a.ResetAtm()
		a.SetState(a.insertCard)
		fmt.Println("All operation will be nil operation:")
		fmt.Println("Error while operation:", err.Error())
	}
}

func NewATM() *ATM {
	atm := &ATM{
		countOFNotes: map[string]int{
			"500": 1000,
			"200": 2000,
			"100": 1000,
		},
		WithdrawAs: &WithdrawNote{},
	}

	atm.insertCard = &InsertCard{
		atm: atm,
	}

	atm.readCard = &ReadCard{
		atm: atm,
	}

	atm.selectAccount = &SelectAccount{
		atm: atm,
	}

	atm.readCard = &ReadCard{
		atm: atm,
	}

	atm.dispenserAmount = &DispenserAmount{
		atm: atm,
	}

	atm.SetState(atm.insertCard)

	return atm
}
