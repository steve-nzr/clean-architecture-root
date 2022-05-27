package domain

import "context"

type MoviesPersistenceWriter interface {
	Create(ctx context.Context, m Movie) error
	Update(ctx context.Context, m Movie) error
}

type MoviesPersistenceReader interface {
	List(ctx context.Context) ([]Movie, error)
	Get(ctx context.Context, id string) (Movie, error)
}
