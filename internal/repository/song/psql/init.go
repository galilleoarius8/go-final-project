package psql

import (
	"context"
	"database/sql"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

type SongPostgres interface {
	Get(ctx context.Context, id int64) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (int64, error)
	GetAllSong(ctx context.Context, offset string, limit string, filter_artist_id string) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]int64, error)
	Update(ctx context.Context, song entity.Song) error
	Delete(ctx context.Context, id int64) error
}

type songConnection struct {
	db *sql.DB
}

func NewSongPostgres(db *sql.DB) SongPostgres {
	return &songConnection{db: db}
}
