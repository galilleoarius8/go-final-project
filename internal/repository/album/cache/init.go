package cache

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/go-redis/redis/v8"
)

type AlbumPostgres interface {
	GetAlbum(ctx context.Context, id int64) (*entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	SetAlbum(ctx context.Context, id int64, Album entity.Album) error
	SetAllAlbum(ctx context.Context, Albums []entity.Album) error
	Delete(ctx context.Context, id int64) error
}

type albumConnection struct {
	client *redis.Client
}

func NewAlbumRedis(cache *redis.Client) AlbumPostgres {
	return &albumConnection{client: cache}
}
