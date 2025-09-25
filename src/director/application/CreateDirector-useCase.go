package application

import (
	"comportaminetoApropiado/src/director/domain"
	"context"
	"errors"
)

type CreateDirectorUseCase struct {
	repo domain.IDirector
}

func NewCreateDirectorUseCase(repo domain.IDirector) *CreateDirectorUseCase {
	return &CreateDirectorUseCase{repo: repo}
}

func (uc *CreateDirectorUseCase) Execute(ctx context.Context, director domain.DirectorUser) (domain.DirectorUser, error) {
	if director.Name == "" || director.Apellidos == "" {
		return domain.DirectorUser{}, errors.New("El nombre y apellidos son obligatorios")
	}
	if director.email == "" || director.password == "" {
		return domain.DirectorUser{}, errors.New("El email y password son obligatorios")
	}
	director.CreatedAt = director.UpdatedAt
	director.ID = director.ID
	createDirector, err := uc.repo.Create(ctx, director)
	if err != nil {
		return domain.DirectorUser{}, err
	}
	return createDirector, nil
}