package request

import domain_users "health-app/domain/users"

type JSONUsers struct {
	UserID   string
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(userReq JSONUsers) domain_users.Users {
	return domain_users.Users{
		UserID:   "",
		Name:     userReq.Name,
		Age:      userReq.Age,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
}