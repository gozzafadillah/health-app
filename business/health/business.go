package business_health

import (
	"errors"
	"fmt"
	domain_health "health-app/domain/health"
	domain_users "health-app/domain/users"
)

type BusinessHealth struct {
	Repo         domain_health.Repository
	UserBusiness domain_users.Business
}

// BMICalculate implements domain_health.Business
func (bh BusinessHealth) BMICalculate(userID string) (float32, error) {
	user, err := bh.Repo.GetDataUser(userID)
	if err != nil {
		return 0, errors.New("user not found")
	}
	fmt.Println("user :", user)
	height := ((float32(user.Height) / 100) * (float32(user.Height) / 100))
	fmt.Println("height :", height)
	BMI := float32(user.Weight) / height
	fmt.Println("BMI :", BMI)
	return BMI, nil
}

// AddDataHealth implements domain_health.Business
func (bh BusinessHealth) AddDataHealth(userID string, domain domain_health.Health) error {
	domain.UserID = userID
	err := bh.Repo.Store(domain)
	if err != nil {
		return errors.New("users failed add data")
	}
	return nil
}

// CalculateIdealWeight implements domain_health.Business
func (bh BusinessHealth) CalculateIdealWeight(userID string) (interface{}, error) {
	user, err := bh.UserBusiness.GetUserByUserID(userID)
	if err != nil {
		return 0, errors.New("user not found")
	}
	weight, err := bh.Repo.GetDataUser(userID)
	if err != nil {
		return 0, errors.New("failed calculate data")
	}
	var idealWeight float32
	if user.Gender == "female" {
		idealWeight = (float32(weight.Weight) - 100) - ((float32(weight.Weight) - 100) * 15 / 100)
	} else {
		idealWeight = (float32(weight.Weight) - 100) - ((float32(weight.Weight) - 100) * 10 / 100)
	}
	BMI, err := bh.BMICalculate(userID)
	if err != nil {
		return 0, errors.New("bmi can't calculate")
	}

	data := map[string]interface{}{
		"userid":       user.UserID,
		"name":         user.Name,
		"gender":       user.Gender,
		"age":          user.Age,
		"weight_ideal": idealWeight,
		"bmi":          BMI,
	}
	return data, err
}

func NewHealthBusiness(repo domain_health.Repository, userBusiness domain_users.Business) domain_health.Business {
	return BusinessHealth{
		Repo:         repo,
		UserBusiness: userBusiness,
	}
}
