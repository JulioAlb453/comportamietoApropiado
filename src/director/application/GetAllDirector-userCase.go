package application

import (
	"comportaminetoApropiado/src/director/domain"
	"context"
)

type GetAllDirectorUserCase struct {
	repo domain.IDirector
}

func NewGetDirectorUserCase(repo domain.IDirector) *GetAllDirectorUserCase {
	return &GetAllDirectorUserCase{repo: repo}
}

func (uc *GetAllDirectorUserCase) Execute(ctx context.Context) ([]domain.DirectorUser, error) {
	return  uc.repo.GetAllDirector(ctx)
}