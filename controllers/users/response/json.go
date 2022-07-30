package response

import domain_users "health-app/domain/users"

type JSONUsers struct {
	UserID   string
	Name     string 
	Age      int    
	Email    string 
}

func FromDomain(users domain_users.Users) JSONUsers {
	return JSONUsers{
		UserID:   users.UserID,
		Name:     users.Name,
		Age:      users.Age,
		Email:    users.Email,
	}
}