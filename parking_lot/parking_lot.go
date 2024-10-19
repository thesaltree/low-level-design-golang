package main

import (
	"fmt"
	"lld_go_parking_lot/vehicles"
	"sync"
	"time"
)

var (
	parkingLotInstance *ParkingLot
	once               sync.Once
)

type ParkingLot struct {
	Name   string
	floors []*ParkingFloor
}

func GetParkingLotInstance() *ParkingLot {

	once.Do(func() {
		parkingLotInstance = &ParkingLot{}
	})
	return parkingLotInstance
}

func (p *ParkingLot) AddFloor(floorID int) {
	p.floors = append(p.floors, NewParkingFloor(floorID))
}

func (p *ParkingLot) DisplayAvailability() {
	fmt.Printf("Parking Lot: %s\n", p.Name)

	for _, floor := range p.floors {
		floor.DisplayFloorStatus(floor)
	}
}

func (p *ParkingLot) findParkingSpot(vehicleType vehicles.VehicleType) (*ParkingSpot, error) {
	for _, floor := range p.floors {
		if spot := floor.FindParkingSpot(vehicleType); spot != nil {
			return spot, nil
		}
	}

	return nil, fmt.Errorf("no available parking spot found for %s", vehicleType)
}

func (p *ParkingLot) ParkVehicle(vehicle vehicles.VehicleInterface) (*ParkingTicket, error) {
	parkingSpot, err := p.findParkingSpot(vehicle.GetVehicleType())
	if err != nil {
		return nil, err
	}

	err = parkingSpot.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	parkingTicket := NewParkingTicket(vehicle, parkingSpot)

	return parkingTicket, nil
}

func (p *ParkingLot) UnparkVehicle(parkingTicket *ParkingTicket) error {
	parkingTicket.SetExitTime(time.Now())
	charge := parkingTicket.CalculateTotalCharge()

	paymentSystem := NewPaymentSystem(charge, parkingTicket)

	if err := paymentSystem.ProcessPayment(); err != nil {
		return fmt.Errorf("payment failed: %v. Vehicle is still parked", err)
	}

	parkingTicket.Spot.RemoveVehicle()

	return nil
}
