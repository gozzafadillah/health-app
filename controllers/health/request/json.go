package request

import domain_health "health-app/domain/health"

type JSONHealth struct {
	Weight int `json:"weight" form:"weight"`
	Height int `json:"height" form:"height"`
	UserID string
}

func ToDomain(healthReq JSONHealth) domain_health.Health {
	return domain_health.Health{
		Weight: healthReq.Weight,
		Height: healthReq.Height,
		UserID: "",
	}
}
