package requests

import "alterra/business/karyawans"

type KaryawanLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(login KaryawanLogin) karyawans.Domain {
	return karyawans.Domain{
		Password: login.Password,
		Email:    login.Password,
	}
}
