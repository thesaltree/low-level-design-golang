package main

import (
	"fmt"
)

type Status string

const (
	PaymentStatusCompleted Status = "Completed"
	PaymentStatusFailed    Status = "Failed"
	PaymentStatusPending   Status = "Pending"
)

type PaymentSystem struct {
	Status        Status
	Amount        float64
	ParkingTicket *ParkingTicket
}

func NewPaymentSystem(amount float64, ticket *ParkingTicket) *PaymentSystem {
	return &PaymentSystem{Status: PaymentStatusPending, Amount: amount, ParkingTicket: ticket}
}

func (p *PaymentSystem) ProcessPayment() error {
	if p.ParkingTicket == nil {
		return fmt.Errorf("payment failed: no parking ticket found")
	}
	if p.ParkingTicket.TotalCharge < p.Amount {
		p.Status = PaymentStatusFailed
		return fmt.Errorf("payment failed: insufficient funds")
	}

	p.Status = PaymentStatusCompleted
	return nil
}

func (p *PaymentSystem) GetPaymentStatus() Status {
	return p.Status
}
