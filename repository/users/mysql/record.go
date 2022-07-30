package repository_users

import (
	domain_users "health-app/domain/users"

	"gorm.io/gorm"
)

type Users struct {
	*gorm.Model
	UserID string
	Name string
	Email string
	Password string
	Age int
}

func ToDomain(rec Users) domain_users.Users{
	return domain_users.Users{
		UserID:   rec.UserID,
		Name:     rec.Name,
		Age:      rec.Age,
		Email:    rec.Email,
		Password: rec.Password,
	}
}