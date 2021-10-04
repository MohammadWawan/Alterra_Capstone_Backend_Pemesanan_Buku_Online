package karyawans

import (
	"alterra/app/middlewares"
	"alterra/business"
	"context"
	"errors"
	"log"
	"time"
)

type KarwayanUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewKaryawanUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) *KarwayanUsecase {
	return &KarwayanUsecase{
		Repo:           repo,
		contextTimeout: timeout,
		jwtAuth:        jwtAuth,
	}
}

func (uc *KarwayanUsecase) Login(ctx context.Context, email string, password string) (Domain, string, error) {
	if email == "" {
		return Domain{}, "", errors.New("email empty")
	}

	if password == "" {
		return Domain{}, "", errors.New("password empty")
	}

	karyawan, err := uc.Repo.GetByEmail(ctx, email)

	if err != nil {
		return Domain{}, "", err
	}

	token, errToken := uc.jwtAuth.GenerateTokenJWT(karyawan.Id)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", errors.New("authentication failed: invalid user credentials")
	}
	return karyawan, token, nil
}

func (uc *KarwayanUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email kosong")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password kosong")
	}

	data, err := uc.Repo.GetByEmail(ctx, domain.Email)
	if data.Id > 0 {
		return Domain{}, errors.New("email has been used")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password is required")
	}
	karyawan, err := uc.Repo.Register(ctx, &domain)

	if err != nil {
		return Domain{}, err
	}
	if err != nil {
		return Domain{}, err
	}

	return karyawan, nil
}

func (uc *KarwayanUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *KarwayanUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	user, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return user, nil
}

func (uc *KarwayanUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	user, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *KarwayanUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
