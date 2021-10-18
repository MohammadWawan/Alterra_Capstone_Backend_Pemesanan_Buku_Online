package responses

import "alterra/business/karyawans"

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
