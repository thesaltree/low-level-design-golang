package atm

import "atm-machine/model"

const (
	InvalidCard = "invalid"
)

// This will be actuall implementation of card reader // i am returning a card from here
type CardReader struct {
}

func NewCardReader() CardReader {
	return CardReader{}
}

// for now it will return some dummy card depend of input for acting as multiple car
func (c CardReader) ReadCard(cardType string) model.Card {
	switch cardType {
	case InvalidCard:
		return model.Card{
			Status: model.NoCard,
		}
	default:
		return model.Card{
			BankName:  "hdfc",
			CardNo:    "34567-sd55498459",
			AccountNo: "bshy4859-sdhhj66",
			Status:    model.Active,
			UserName:  "Himanshu sharma"}

	}

}
