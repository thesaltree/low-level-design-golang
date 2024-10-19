package main

import (
	"fmt"
	"lld_go_parking_lot/vehicles"
	"sync"
)

type ParkingSpot struct {
	SpotID         int
	VehicleType    vehicles.VehicleType
	CurrentVehicle *vehicles.VehicleInterface
	lock           sync.Mutex
}

func NewParkingSpot(spotID int, vehicleType vehicles.VehicleType) *ParkingSpot {
	return &ParkingSpot{SpotID: spotID, VehicleType: vehicleType}
}

func (p *ParkingSpot) IsParkingSpotFree() bool {
	return p.CurrentVehicle == nil
}

func (p *ParkingSpot) ParkVehicle(vehicle vehicles.VehicleInterface) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if vehicle.GetVehicleType() != p.VehicleType {
		return fmt.Errorf("vehicle type mismatch: expected %s, got %s", p.VehicleType, vehicle.GetVehicleType())
	}
	if p.CurrentVehicle != nil {
		return fmt.Errorf("parking spot already occupied")
	}

	p.CurrentVehicle = &vehicle
	return nil
}

func (p *ParkingSpot) RemoveVehicle() {
	p.CurrentVehicle = nil
}
