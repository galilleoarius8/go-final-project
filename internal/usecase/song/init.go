package usecase

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	songRepository "github.com/galileoarius8/final-project-edspert/internal/repository/song"
)

type SongUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (*entity.Song, error)
	GetAllSong(ctx context.Context, offset string, limit string, filter_album_id string) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]entity.Song, error)
	Update(ctx context.Context, song entity.Song) (entity.Song, error)
	Delete(ctx context.Context, id int64) error
}

type songUsecase struct {
	songRepository songRepository.SongRepository
}

func NewSongUsecase(songRepository songRepository.SongRepository) SongUsecase {
	return &songUsecase{
		songRepository: songRepository,
	}
}
