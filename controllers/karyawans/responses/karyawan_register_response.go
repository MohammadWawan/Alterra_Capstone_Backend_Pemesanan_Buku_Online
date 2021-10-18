package responses

import (
	"alterra/business/karyawans"
	"time"
)

type KaryawanRegisterResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromKaryawansRegisterDomain(domain karyawans.Domain) KaryawanResponse {
	return KaryawanResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
