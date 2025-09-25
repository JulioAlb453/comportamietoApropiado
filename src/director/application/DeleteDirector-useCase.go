package application

import (
	"comportaminetoApropiado/src/director/domain"
	"context"
)

type DeleteDirectorUseCase struct {
	repo domain.IDirector
}

func NewDeleteDirectorUserCase(repo domain.IDirector) *DeleteDirectorUseCase{
	return &DeleteDirectorUseCase{repo: repo}
}

func (uc *DeleteDirectorUseCase) Execute(ctx context.Context, id string ) (domain.DirectorUser, error){
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return domain.DirectorUser{}, err
	}
	return domain.DirectorUser{}, nil
}
