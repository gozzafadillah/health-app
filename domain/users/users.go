package domain_users

type Users struct {
	UserID   string
	Name     string
	Age      int
	Email    string
	Password string
}

type Business interface {
	Register(domain Users) error
	Login(email, password string) (string, error)
	GetUserByUserID(userID string)(Users, error)
}

type Repository interface {
	Store(domain Users) error
	GetUserByEmail(email string) ( Users, error)
	GetUserByUserID(userID string)(Users, error)
	Auth(email, password string) error
}