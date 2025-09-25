package application

import (
	"comportaminetoApropiado/src/director/domain"
	"context"
)

type GetDirectorByIDUseCase struct {
	repo domain.IDirector
}

func NewGetDirectorByIDUseCase(repo domain.IDirector) *GetDirectorByIDUseCase {
	return &GetDirectorByIDUseCase{repo: repo}
} 

func (uc *GetDirectorByIDUseCase) Execute(ctx context.Context, id string) (domain.DirectorUser, error) {
	director, err := uc.repo.GetDirectorByID(ctx, id)
	if err != nil {
		return domain.DirectorUser{}, err
	}
	return director, nil
}
