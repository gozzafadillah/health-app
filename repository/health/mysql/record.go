package repository_health

import (
	domain_health "health-app/domain/health"

	"gorm.io/gorm"
)

type Health struct {
	*gorm.Model
	Weight int
	Height int
	UserID string `gorm:"unique"`
}

func ToDomain(rec Health) domain_health.Health {
	return domain_health.Health{
		Weight: rec.Weight,
		Height: rec.Height,
		UserID: rec.UserID,
	}
}
