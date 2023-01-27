package main

import (
	"fmt"
	"log"
	"os"

	"github.com/galileoarius8/final-project-edspert/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		// Set to release mode
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Load redis
	cache, err := config.OpenCache(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	// Init clean arch
	repository_artist := config.InitRepository_artist(db, cache)
	usecase_artist := config.InitUsecase_artist(repository_artist.ArtistRepository)
	handler_artist := config.InitHandler_artist(usecase_artist.ArtistUsecase)
	repository_album := config.InitRepository_album(db, cache)
	usecase_album := config.InitUsecase_album(repository_album.AlbumRepository)
	handler_album := config.InitHandler_album(usecase_album.AlbumUsecase)
	repository_song := config.InitRepository_song(db, cache)
	usecase_song := config.InitUsecase_song(repository_song.SongRepository)
	handler_song := config.InitHandler_song(usecase_song.SongUsecase)

	// Create the API
	artistRoutes := r.Group("/api/v1/artists")
	{
		artistRoutes.GET("/", handler_artist.ArtistHandler.GetAllArtist)
		artistRoutes.POST("/", handler_artist.ArtistHandler.Create)
		artistRoutes.POST("/batch", handler_artist.ArtistHandler.BatchCreate)
		artistRoutes.GET("/:id", handler_artist.ArtistHandler.Get)
		artistRoutes.PUT("/:id", handler_artist.ArtistHandler.Update)
		artistRoutes.DELETE("/:id", handler_artist.ArtistHandler.Delete)
	}

	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", handler_album.AlbumHandler.GetAllAlbum)
		albumRoutes.POST("/", handler_album.AlbumHandler.Create)
		albumRoutes.POST("/batch", handler_album.AlbumHandler.BatchCreate)
		albumRoutes.GET("/:id", handler_album.AlbumHandler.Get)
		albumRoutes.PUT("/:id", handler_album.AlbumHandler.Update)
		albumRoutes.DELETE("/:id", handler_album.AlbumHandler.Delete)
	}

	songRoutes := r.Group("/api/v1/songs")
	{
		songRoutes.GET("/", handler_song.SongHandler.GetAllSong)
		songRoutes.POST("/", handler_song.SongHandler.Create)
		songRoutes.POST("/batch", handler_song.SongHandler.BatchCreate)
		songRoutes.GET("/:id", handler_song.SongHandler.Get)
		songRoutes.PUT("/:id", handler_song.SongHandler.Update)
		songRoutes.DELETE("/:id", handler_song.SongHandler.Delete)
	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
