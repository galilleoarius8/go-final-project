package repository

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

func (repo *songRepository) Create(ctx context.Context, song *entity.Song) (int64, error) {
	return repo.postgres.Create(ctx, song)
}

// It will call the function Get in psql/song
func (repo *songRepository) Get(ctx context.Context, id int64) (*entity.Song, error) {
	return repo.postgres.Get(ctx, id)
}

// It will call the function GetAllsong in psql/song
func (repo *songRepository) GetAllSong(ctx context.Context, offset string, limit string, filter_album_id string) ([]entity.Song, error) {
	return repo.postgres.GetAllSong(ctx, offset, limit, filter_album_id)
}

// It will call the function BatchCreate in psql/song
func (repo *songRepository) BatchCreate(ctx context.Context, songs []entity.Song) ([]int64, error) {
	return repo.postgres.BatchCreate(ctx, songs)
}

// It will call the function Update in psql/song
func (repo *songRepository) Update(ctx context.Context, song entity.Song) error {
	return repo.postgres.Update(ctx, song)
}

// It will call the function Delete in psql/song
func (repo *songRepository) Delete(ctx context.Context, id int64) error {
	return repo.postgres.Delete(ctx, id)
}

func (repo *songRepository) GetSongCache(ctx context.Context, id int64) (*entity.Song, error) {
	return repo.cache.GetSong(ctx, id)
}

func (repo *songRepository) GetAllSongCache(ctx context.Context) ([]entity.Song, error) {
	return repo.cache.GetAllSong(ctx)
}

func (repo *songRepository) SetSongCache(ctx context.Context, id int64, song entity.Song) error {
	return repo.cache.SetSong(ctx, id, song)
}

func (repo *songRepository) SetAllSongCache(ctx context.Context, songs []entity.Song) error {
	return repo.cache.SetAllSong(ctx, songs)
}

func (repo *songRepository) DeleteSongCache(ctx context.Context, id int64) error {
	return repo.cache.Delete(ctx, id)
}
