package business_users

import (
	"errors"
	middlewares "health-app/app/middlewares"
	domain_users "health-app/domain/users"

	"github.com/google/uuid"
)

type BusinessUsers struct {
	Repo   domain_users.Repository
	JwtCon *middlewares.ConfigJwt
}

// GetUserByUserID implements domain_users.Business
func (bu BusinessUsers) GetUserByUserID(userID string) (domain_users.Users, error) {
	data, err := bu.Repo.GetUserByUserID(userID)
	if err != nil {
		return domain_users.Users{}, errors.New("user not found")
	}
	return data, nil
}

// Login implements domain_users.Business
func (bu BusinessUsers) Login(email string, password string) (string, error) {
	data, err := bu.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	err = bu.Repo.Auth(email, password)
	if err != nil {
		return "", errors.New("email and password miss match")
	}
	token, err := bu.JwtCon.GenerateToken(data.UserID)
	if err != nil {
		return "", errors.New("token not generate")
	}
	return token, nil
}

// Register implements domain_users.Business
func (bu BusinessUsers) Register(domain domain_users.Users) error {
	domain.UserID = uuid.New().String()
	err := bu.Repo.Store(domain)
	if err != nil {
		return errors.New("register failed")
	}
	return nil
}

func NewUsersBusiness(repo domain_users.Repository, jwt *middlewares.ConfigJwt) domain_users.Business {
	return BusinessUsers{
		Repo:   repo,
		JwtCon: jwt,
	}
}
