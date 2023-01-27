package psql

import (
	"context"
	"database/sql"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

type AlbumPostgres interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (int64, error)
	GetAllAlbum(ctx context.Context, offset string, limit string, filter_artist_id string) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error)
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id int64) error
}

type albumConnection struct {
	db *sql.DB
}

func NewAlbumPostgres(db *sql.DB) AlbumPostgres {
	return &albumConnection{db: db}
}
