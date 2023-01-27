package config

import (
	albumRepository "github.com/galileoarius8/final-project-edspert/internal/repository/album"
	artistRepository "github.com/galileoarius8/final-project-edspert/internal/repository/artists"
	songRepository "github.com/galileoarius8/final-project-edspert/internal/repository/song"
	albumUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/album"
	artistUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/artists"
	songUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/song"
)

type Usecase_artist struct {
	ArtistUsecase artistUsecase.ArtistUsecase
}

// Function to initialize usecase
func InitUsecase_artist(artistRepository artistRepository.ArtistRepository) Usecase_artist {
	return Usecase_artist{
		ArtistUsecase: artistUsecase.NewArtistUsecase(artistRepository),
	}
}

type Usecase_album struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

// Function to initialize usecase
func InitUsecase_album(albumRepository albumRepository.AlbumRepository) Usecase_album {
	return Usecase_album{
		AlbumUsecase: albumUsecase.NewAlbumUsecase(albumRepository),
	}
}

type Usecase_song struct {
	SongUsecase songUsecase.SongUsecase
}

// Function to initialize usecase
func InitUsecase_song(songRepository songRepository.SongRepository) Usecase_song {
	return Usecase_song{
		SongUsecase: songUsecase.NewSongUsecase(songRepository),
	}
}
