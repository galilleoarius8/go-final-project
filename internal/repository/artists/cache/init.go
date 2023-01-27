package cache

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/go-redis/redis/v8"
)

type ArtistPostgres interface {
	GetArtist(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	SetArtist(ctx context.Context, id int64, Artist entity.Artist) error
	SetAllArtist(ctx context.Context, Artists []entity.Artist) error
	Delete(ctx context.Context, id int64) error
}

type artistConnection struct {
	client *redis.Client
}

func NewArtistRedis(cache *redis.Client) ArtistPostgres {
	return &artistConnection{client: cache}
}
