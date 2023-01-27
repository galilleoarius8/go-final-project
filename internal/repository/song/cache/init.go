package cache

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/go-redis/redis/v8"
)

type SongPostgres interface {
	GetSong(ctx context.Context, id int64) (*entity.Song, error)
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	SetSong(ctx context.Context, id int64, Song entity.Song) error
	SetAllSong(ctx context.Context, Songs []entity.Song) error
	Delete(ctx context.Context, id int64) error
}

type songConnection struct {
	client *redis.Client
}

func NewSongRedis(cache *redis.Client) SongPostgres {
	return &songConnection{client: cache}
}
