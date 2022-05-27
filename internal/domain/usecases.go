package domain

import "context"

type MovieCreator interface {
	Create(ctx context.Context, input CreateMovie) (EntityID, error)
}
