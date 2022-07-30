package business_health

import (
	"errors"
	domain_health "health-app/domain/health"
	domain_users "health-app/domain/users"
)

type BusinessHealth struct {
	Repo         domain_health.Repository
	UserBusiness domain_users.Business
}

// AddDataHealth implements domain_health.Business
func (bh BusinessHealth) AddDataHealth(userID string,domain domain_health.Health) error {
	domain.UserID = userID
	err := bh.Repo.Store(domain)
	if err != nil {
		return errors.New("users failed add data")
	}
	return nil
}

// CalculateIdealWeight implements domain_health.Business
func (bh BusinessHealth) CalculateIdealWeight(userID string) (interface{}, error) {
	weight, err := bh.Repo.GetDataUser(userID)
	if err != nil {
		return 0, errors.New("failed calculate data")
	}
	idealWeight := (float32(weight.Weight) - 100) - ((float32(weight.Weight) -100) * 10/100)
	user, err := bh.UserBusiness.GetUserByUserID(userID)

	data := map[string]interface{}{
		"userid": user.UserID,
		"name": user.Name,
		"age": user.Age,
		"weight_ideal": idealWeight,
	}
	return data, err
}

func NewHealthBusiness(repo domain_health.Repository, userBusiness domain_users.Business) domain_health.Business {
	return BusinessHealth{
		Repo:         repo,
		UserBusiness: userBusiness,
	}
}
