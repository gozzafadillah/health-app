package repository_users

import (
	domain_users "health-app/domain/users"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

// GetUserByEmail implements domain_users.Repository
func (ur UsersRepo) GetUserByEmail(email string) (domain_users.Users, error) {
	rec := Users{}
	err := ur.DB.Where("email = ?", email).First(&rec).Error
	return ToDomain(rec), err
}

// Auth implements domain_users.Repository
func (ur UsersRepo) Auth(email string, password string) error {
	rec := Users{}
	err := ur.DB.Where("email = ? AND password = ?", email, password).First(&rec).Error
	return err
}

// Store implements domain_users.Repository
func (ur UsersRepo) Store(domain domain_users.Users) error {
	err := ur.DB.Create(domain).Error
	return err
}

func NewUsersRepo(db *gorm.DB) domain_users.Repository {
	return UsersRepo{
		DB: db,
	}
}
