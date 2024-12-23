package service

import (
	"app/internal"
	"errors"
	"strings"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	count := 0.0
	sum := 0.0
	v, err := s.rp.FindAll()

	for _, vehicle := range v {
		if strings.ToLower(vehicle.Brand) == brand {
			count++
			sum += float64(vehicle.MaxSpeed)
		}
	}

	if count == 0.0 {
		err = errors.New("brand not found")
		v = nil
		return
	}

	averageSpeed = sum / count

	return
}
