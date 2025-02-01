package services

import (
	"time"

	"gorm.io/gorm"
	"rental.com/m/internal/models"
)

type VehicleService struct {
	db *gorm.DB
}

func NewVehicleService(db *gorm.DB) *VehicleService {
	return &VehicleService{db: db}
}
func (s *VehicleService) CreateVehicle(vehicle *models.Vehicle) error {
	return s.db.Create(vehicle).Error
}
func (s *VehicleService) GetAvailableVehicles(startDate, endDate time.Time) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := s.db.Where("status = ? AND id NOT IN (SELECT vehicle_id FROM bookings WHERE ? < end_date AND ? > start_date)",
		"available", startDate, endDate).Find(&vehicles).Error
	return vehicles, err
}
