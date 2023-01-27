package usecase

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	albumRepository "github.com/galileoarius8/final-project-edspert/internal/repository/album"
)

type AlbumUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (*entity.Album, error)
	GetAllAlbum(ctx context.Context, offset string, limit string, filter_artist_id string) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]entity.Album, error)
	Update(ctx context.Context, album entity.Album) (entity.Album, error)
	Delete(ctx context.Context, id int64) error
}

type albumUsecase struct {
	albumRepository albumRepository.AlbumRepository
}

func NewAlbumUsecase(albumRepository albumRepository.AlbumRepository) AlbumUsecase {
	return &albumUsecase{
		albumRepository: albumRepository,
	}
}
