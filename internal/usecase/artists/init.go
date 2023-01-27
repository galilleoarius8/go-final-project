package usecase

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	artistRepository "github.com/galileoarius8/final-project-edspert/internal/repository/artists"
)

type ArtistUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error)
	GetAllArtist(ctx context.Context, offset string, limit string) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error)
	Update(ctx context.Context, artist entity.Artist) (entity.Artist, error)
	Delete(ctx context.Context, id int64) error
}

type artistUsecase struct {
	artistRepository artistRepository.ArtistRepository
}

func NewArtistUsecase(artistRepository artistRepository.ArtistRepository) ArtistUsecase {
	return &artistUsecase{
		artistRepository: artistRepository,
	}
}
