package config

import (
	albumHandler "github.com/galileoarius8/final-project-edspert/internal/handler/album"
	albumUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/album"

	artistHandler "github.com/galileoarius8/final-project-edspert/internal/handler/artists"
	artistUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/artists"

	songHandler "github.com/galileoarius8/final-project-edspert/internal/handler/song"
	songUsecase "github.com/galileoarius8/final-project-edspert/internal/usecase/song"
)

type Handler_artist struct {
	ArtistHandler artistHandler.ArtistHandler
}

// Function to initialize handler
func InitHandler_artist(artistUsecase artistUsecase.ArtistUsecase) Handler_artist {
	return Handler_artist{
		ArtistHandler: artistHandler.NewArtistHandler(artistUsecase),
	}
}

type Handler_album struct {
	AlbumHandler albumHandler.AlbumHandler
}

// Function to initialize handler
func InitHandler_album(albumUsecase albumUsecase.AlbumUsecase) Handler_album {
	return Handler_album{
		AlbumHandler: albumHandler.NewAlbumHandler(albumUsecase),
	}
}

type Handler_song struct {
	SongHandler songHandler.SongHandler
}

// Function to initialize handler
func InitHandler_song(songUsecase songUsecase.SongUsecase) Handler_song {
	return Handler_song{
		SongHandler: songHandler.NewSongHandler(songUsecase),
	}
}
