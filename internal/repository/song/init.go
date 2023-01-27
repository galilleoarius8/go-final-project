package repository

import (
	"context"
	"database/sql"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/galileoarius8/final-project-edspert/internal/repository/song/cache"
	"github.com/galileoarius8/final-project-edspert/internal/repository/song/psql"
	"github.com/go-redis/redis/v8"
)

type SongRepository interface {
	Get(ctx context.Context, id int64) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (int64, error)
	GetAllSong(ctx context.Context, offset string, limit string, filter_album_id string) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]int64, error)
	Update(ctx context.Context, song entity.Song) error
	Delete(ctx context.Context, id int64) error

	GetSongCache(ctx context.Context, id int64) (*entity.Song, error)
	GetAllSongCache(ctx context.Context) ([]entity.Song, error)
	SetSongCache(ctx context.Context, id int64, song entity.Song) error
	SetAllSongCache(ctx context.Context, songs []entity.Song) error
	DeleteSongCache(ctx context.Context, id int64) error
}

type songRepository struct {
	postgres psql.SongPostgres
	cache    cache.SongPostgres
}

func NewSongRepository(db *sql.DB, client *redis.Client) SongRepository {
	return &songRepository{
		postgres: psql.NewSongPostgres(db),
		cache:    cache.NewSongRedis(client),
	}
}
