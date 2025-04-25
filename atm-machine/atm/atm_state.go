package atm

type ATMState interface {
	StateName() string
	InsertCard() error
	InsertPin() error
	SelectAccount() error
	GetCardDetail() error
	DispenserAmount() error
}

type ATMAbstract struct{}

func (s *ATMAbstract) InsertCard() error {
	return nil
}

func (s *ATMAbstract) InsertPin() error {
	return nil
}

func (s *ATMAbstract) AuthticateCard() error {
	return nil
}

func (s *ATMAbstract) DispenserAmount() error {
	return nil
}

func (s *ATMAbstract) SelectAccount() error {
	return nil

}

func (s *ATMAbstract) GetCardDetail() error {
	return nil

}

func (s *ATMAbstract) StateName() string {
	return "ATMAbstract"
}
