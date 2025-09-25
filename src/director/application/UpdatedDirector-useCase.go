package application

import (
	"comportaminetoApropiado/src/director/domain"
	"context"
	"errors"
)

type UpdateDirectorUseCase struct {
	repo domain.IDirector
}

func NewUpdateDirectorUseCase(repo domain.IDirector) *UpdateDirectorUseCase {
	return &UpdateDirectorUseCase{repo: repo}
}

func (uc *UpdateDirectorUseCase) Execute(ctx context.Context, director domain.DirectorUser) (domain.DirectorUser, error) {
	if domain.Name == "" || domain.Apellidos == "" {
		return domain.DirectorUser{}, errors.New("El nombre y apellidos son obligatorios")
	}
	if domain.email == "" || domain.password == "" {
		return domain.DirectorUser{}, errors.New("El email y password son obligatorios")
	}
	existingDirector, err := uc.repo.GetDirectorByID(ctx, director.ID)
	if err != nil {
		return domain.DirectorUser{}, err
	}
	existingDirector.Name = director.Name
	existingDirector.Apellidos = director.Apellidos
	existingDirector.email = director.email
	existingDirector.password = director.password
	existingDirector.UpdatedAt = director.UpdatedAt

	updateDirector, err := uc.repo.Update(ctx, existingDirector)

	if err != nil {
		return domain.DirectorUser{}, err
	}
	return updateDirector, nil

}
