package responses

import (
	"alterra/business/karyawans"
	"time"
)

type KaryawanResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type LoginResponse struct {
	SessionToken string
	Karyawan     interface{}
}

func FromDomain(domain karyawans.Domain) KaryawanResponse {
	return KaryawanResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromKaryawansListDomain(domain []karyawans.Domain) []KaryawanResponse {
	var list []KaryawanResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

func FromKaryawansDomainToLogin(domain karyawans.Domain, token string) LoginResponse {
	response := KaryawanResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}

	loginResponse := LoginResponse{}
	loginResponse.SessionToken = token
	loginResponse.Karyawan = response
	return loginResponse
}

func FromKaryawansDomain(domain karyawans.Domain) KaryawanResponse {
	return KaryawanResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
