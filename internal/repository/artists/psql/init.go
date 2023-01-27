package psql

import (
	"context"
	"database/sql"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

//==========================================================================ARTIST

type ArtistPostgres interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (int64, error)
	GetAllArtist(ctx context.Context, offset string, limit string) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error)
	Update(ctx context.Context, artist entity.Artist) error
	Delete(ctx context.Context, id int64) error
}

type artistConnection struct {
	db *sql.DB
}

func NewArtistPostgres(db *sql.DB) ArtistPostgres {
	return &artistConnection{db: db}
}

//======================================================================ALBUM
