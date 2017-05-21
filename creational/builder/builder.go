package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	//Implementation goes here
}

func (f *ManufacturingDirector) Construct() {
	//Implementation goes here
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}

type BikeBuilder struct{}

func (b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetSturcture() BuildProcess {
	return nil
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}

type MotorBikeBuilder struct{}

func (m *MotorBikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (m *MotorBikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (m *MotorBikeBuilder) SetStructure() BuildProcess {
	return nil
}

func (m *MotorBikeBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}