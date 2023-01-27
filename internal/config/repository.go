package config

import (
	"database/sql"

	albumRepository "github.com/galileoarius8/final-project-edspert/internal/repository/album"
	artistRepository "github.com/galileoarius8/final-project-edspert/internal/repository/artists"
	songRepository "github.com/galileoarius8/final-project-edspert/internal/repository/song"
	"github.com/go-redis/redis/v8"
)

type Repository_artist struct {
	ArtistRepository artistRepository.ArtistRepository
}

// Function to initialize repository
func InitRepository_artist(db *sql.DB, cache *redis.Client) Repository_artist {
	return Repository_artist{
		ArtistRepository: artistRepository.NewArtistRepository(db, cache),
	}
}

type Repository_album struct {
	AlbumRepository albumRepository.AlbumRepository
}

// Function to initialize repository
func InitRepository_album(db *sql.DB, cache *redis.Client) Repository_album {
	return Repository_album{
		AlbumRepository: albumRepository.NewAlbumRepository(db, cache),
	}
}

type Repository_song struct {
	SongRepository songRepository.SongRepository
}

// Function to initialize repository
func InitRepository_song(db *sql.DB, cache *redis.Client) Repository_song {
	return Repository_song{
		SongRepository: songRepository.NewSongRepository(db, cache),
	}
}
