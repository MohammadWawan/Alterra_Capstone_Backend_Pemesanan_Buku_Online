package requests

import "alterra/business/karyawans"

type KaryawanRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *KaryawanRegister) ToDomain() *karyawans.Domain {
	return &karyawans.Domain{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
