package main

import (
	"fmt"
	"lld_go_parking_lot/vehicles"
)

const (
	CarSpotCount        = 5
	VanSpotCount        = 3
	TruckSpotCount      = 2
	MotorcycleSpotCount = 10
)

type ParkingFloor struct {
	FloorID      int
	ParkingSpots map[vehicles.VehicleType]map[int]*ParkingSpot
}

func NewParkingFloor(floorID int) *ParkingFloor {
	parkingSpots := make(map[vehicles.VehicleType]map[int]*ParkingSpot)

	parkingSpots[vehicles.CarType] = createParkingSpots(CarSpotCount, vehicles.CarType)
	parkingSpots[vehicles.VanType] = createParkingSpots(VanSpotCount, vehicles.VanType)
	parkingSpots[vehicles.TruckType] = createParkingSpots(TruckSpotCount, vehicles.TruckType)
	parkingSpots[vehicles.MotorcycleType] = createParkingSpots(MotorcycleSpotCount, vehicles.MotorcycleType)

	return &ParkingFloor{FloorID: floorID, ParkingSpots: parkingSpots}
}

func createParkingSpots(count int, vehicleType vehicles.VehicleType) map[int]*ParkingSpot {
	spots := make(map[int]*ParkingSpot)
	for i := 1; i <= count; i++ {
		spots[i] = NewParkingSpot(i, vehicleType)
	}
	return spots
}

func (p *ParkingFloor) FindParkingSpot(vehicleType vehicles.VehicleType) *ParkingSpot {
	for _, spot := range p.ParkingSpots[vehicleType] {
		if spot.IsParkingSpotFree() {
			return spot
		}
	}
	return nil
}

func (p *ParkingFloor) DisplayFloorStatus(parkingFloor *ParkingFloor) {
	fmt.Printf("Floor ID: %d\n", parkingFloor.FloorID)

	for vehicleType, spotMap := range parkingFloor.ParkingSpots {
		fmt.Printf("\n%s Spots:\n", vehicleType)
		count := 0

		for _, spot := range spotMap {
			if spot.IsParkingSpotFree() {
				count++
			}
		}

		fmt.Printf("\n%s Spot: %d Free\n", vehicleType, count)
	}
}
