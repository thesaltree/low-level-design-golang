package model

type cardStatus string

const (
	Active  cardStatus = "ACTIVE"
	Expiry  cardStatus = "EXPIRE"
	Blocked cardStatus = "BLOCKED"
	NoCard  cardStatus = "NOCARD"
)

type Card struct {
	BankName  string
	CardNo    string
	AccountNo string
	Status    cardStatus
	UserName  string
}
