package vehicles

type VehicleType string

const (
	CarType        VehicleType = "Car"
	VanType        VehicleType = "Van"
	TruckType      VehicleType = "Truck"
	MotorcycleType VehicleType = "Motorcycle"
)

var vehicleCosts = map[VehicleType]float64{
	CarType:        100,
	VanType:        200,
	TruckType:      300,
	MotorcycleType: 50,
}

type Vehicle struct {
	LicenceNumber string
	VehicleType   VehicleType
	Cost          float64
}

type VehicleInterface interface {
	GetLicenceNumber() string
	GetVehicleType() VehicleType
	GetVehicleCost() float64
}

func (v *Vehicle) GetLicenceNumber() string {
	return v.LicenceNumber
}

func (v *Vehicle) GetVehicleType() VehicleType {
	return v.VehicleType
}

func (v *Vehicle) GetVehicleCost() float64 {
	return v.Cost
}

func NewVehicle(licenceNumber string, vehicleType VehicleType) *Vehicle {
	cost := vehicleCosts[vehicleType]
	return &Vehicle{LicenceNumber: licenceNumber, VehicleType: vehicleType, Cost: cost}
}

type Car struct {
	Vehicle
}

func NewCar(licenceNumber string) *Car {
	return &Car{Vehicle: *NewVehicle(licenceNumber, CarType)}
}

type Van struct {
	Vehicle
}

func NewVan(licenceNumber string) *Van {
	return &Van{Vehicle: *NewVehicle(licenceNumber, VanType)}
}

type Truck struct {
	Vehicle
}

func NewTruck(licenceNumber string) *Truck {
	return &Truck{Vehicle: *NewVehicle(licenceNumber, TruckType)}
}

type Motorcycle struct {
	Vehicle
}

func NewMotorcycle(licenceNumber string) *Motorcycle {
	return &Motorcycle{Vehicle: *NewVehicle(licenceNumber, MotorcycleType)}
}
