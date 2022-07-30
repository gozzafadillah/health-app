package repository_health

import (
	domain_health "health-app/domain/health"

	"gorm.io/gorm"
)

type HealthRepo struct {
	DB *gorm.DB
}

// Store implements domain_health.Repository
func (hr HealthRepo) Store(domain domain_health.Health) error {
	err := hr.DB.Create(&domain).Error
	return err
}

// GetDataUser implements domain_health.Repository
func (hr HealthRepo) GetDataUser(userID string) (domain_health.Health, error) {
	rec := Health{}
	err := hr.DB.Where("user_id = ? ", userID).First(&rec).Error
	return ToDomain(rec), err
}

func NewHealthRepo(db *gorm.DB) domain_health.Repository {
	return HealthRepo{
		DB: db,
	}
}
