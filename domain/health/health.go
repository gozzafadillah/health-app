package domain_health

type Health struct {
	Weight int
	Height int
	UserID string
}

type Business interface {
	AddDataHealth(userID string, domain Health) error
	CalculateIdealWeight(userID string) (interface{}, error)
}

type Repository interface {
	Store(domain Health) error
	GetDataUser(userID string) (Health, error)
}