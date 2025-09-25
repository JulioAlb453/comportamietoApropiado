package domain

import (
	"context"
)

type IDirector interface {
	GetDirector(ctx context.Context, id int) (*DirectorUser, error)
	GetDirectorById(ctx context.Context, id int) (*DirectorUser, error)
	CreateDirector(ctx context.Context, director *DirectorUser) error
	UpdateDirector(ctx context.Context, director *DirectorUser) error
	DeleteDirector(ctx context.Context, id int) error
}
