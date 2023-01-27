package usecase

import (
	"context"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

func (usecase *artistUsecase) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	// Get from cache
	artist, err := usecase.artistRepository.GetArtistCache(ctx, id)
	if err != nil {
		return artist, err
	}

	if artist.ID != 0 {
		return artist, nil
	}

	// Get from db
	artist, err = usecase.artistRepository.Get(ctx, id)
	if err != nil {
		return artist, err
	}

	// Set to cache
	if err = usecase.artistRepository.SetArtistCache(ctx, id, *artist); err != nil {
		return artist, err
	}

	return artist, nil
}

// It will call the function Create in artist repository
func (usecase *artistUsecase) Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error) {
	var newArtist *entity.Artist

	// Create artist
	id, err := usecase.artistRepository.Create(ctx, artist)
	if err != nil {
		return newArtist, err
	}

	// Find new artist
	newArtist, err = usecase.artistRepository.Get(ctx, id)
	if err != nil {
		return newArtist, err
	}

	// Find all artists
	// artists, err := usecase.artistRepository.GetAllArtist(ctx, offset, limit)
	// if err != nil {
	// 	return newArtist, err
	// }

	// Set to specific cache
	if err = usecase.artistRepository.SetArtistCache(ctx, id, *newArtist); err != nil {
		return newArtist, err
	}

	// Set all cache
	// if err = usecase.artistRepository.SetAllArtistCache(ctx, artists); err != nil {
	// 	return newArtist, err
	// }

	return newArtist, nil
}

// It will call the function GetAllArtist in artist repository
func (usecase *artistUsecase) GetAllArtist(ctx context.Context, offset string, limit string) ([]entity.Artist, error) {
	var artists []entity.Artist

	// // Get from cache
	// artists, err := usecase.artistRepository.GetAllArtistCache(ctx)
	// if err != nil {
	// 	return artists, err
	// }

	// if len(artists) > 0 {
	// 	return artists, nil
	// }

	// Get from db
	artists, err := usecase.artistRepository.GetAllArtist(ctx, offset, limit)
	if err != nil {
		return artists, err
	}

	// // Set to cache
	// if err = usecase.artistRepository.SetAllArtistCache(ctx, artists); err != nil {
	// 	return artists, err
	// }

	return artists, nil
}

// It will call the function BatchCreate in artist repository
func (usecase *artistUsecase) BatchCreate(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error) {
	var newArtists []entity.Artist

	// Batch create and get the new id
	ids, err := usecase.artistRepository.BatchCreate(ctx, artists)
	if err != nil {
		return newArtists, err
	}

	// Get detail artist by ids
	for _, id := range ids {
		// Get from db
		artist, err := usecase.artistRepository.Get(ctx, id)
		if err != nil {
			return newArtists, err
		}

		// Set to specific cache
		if err = usecase.artistRepository.SetArtistCache(ctx, id, *artist); err != nil {
			return newArtists, err
		}

		newArtists = append(newArtists, *artist)
	}

	// Find all artists
	// allArtists, err := usecase.artistRepository.GetAllArtist(ctx)
	// if err != nil {
	// 	return newArtists, err
	// }

	// Set all cache
	// if err = usecase.artistRepository.SetAllArtistCache(ctx, allArtists); err != nil {
	// 	return newArtists, err
	// }

	return newArtists, nil
}

// It will call the function Update in artist repository
func (usecase *artistUsecase) Update(ctx context.Context, artist entity.Artist) (entity.Artist, error) {
	var updatedArtist *entity.Artist

	// Update artist
	err := usecase.artistRepository.Update(ctx, artist)
	if err != nil {
		return *updatedArtist, err
	}

	// Find new artist
	updatedArtist, err = usecase.artistRepository.Get(ctx, artist.ID)
	if err != nil {
		return *updatedArtist, err
	}

	// Find all artists
	// artists, err := usecase.artistRepository.GetAllArtist(ctx)
	// if err != nil {
	// 	return *updatedArtist, err
	// }

	// Set to specific cache
	if err = usecase.artistRepository.SetArtistCache(ctx, artist.ID, *updatedArtist); err != nil {
		return *updatedArtist, err
	}

	// Set all cache
	// if err = usecase.artistRepository.SetAllArtistCache(ctx, artists); err != nil {
	// 	return *updatedArtist, err
	// }

	return *updatedArtist, nil
}

// It will call the function Delete in artist repository
func (usecase *artistUsecase) Delete(ctx context.Context, id int64) error {
	// Delete from db
	if err := usecase.artistRepository.Delete(ctx, id); err != nil {
		return err
	}

	// Delete from cache
	if err := usecase.artistRepository.DeleteArtistCache(ctx, id); err != nil {
		return err
	}

	// Find all artists
	// artists, err := usecase.artistRepository.GetAllArtist(ctx)
	// if err != nil {
	// 	return err
	// }

	// Set all cache
	// if err = usecase.artistRepository.SetAllArtistCache(ctx, artists); err != nil {
	// 	return err
	// }

	return nil
}
