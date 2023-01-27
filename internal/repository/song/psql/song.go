package psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
)

func (repo *songConnection) Create(ctx context.Context, song *entity.Song) (int64, error) {
	// The query insert
	query := `
        INSERT INTO public.song (title,album_id,lyrics) 
        VALUES ($1,$2,$3)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, song.Title, song.Album_id, song.Lyrics).Scan(&song.ID)
	if err != nil {
		return 0, err
	}

	return song.ID, nil
}

// Get is function to get specific song by id from database
func (repo *songConnection) Get(ctx context.Context, id int64) (*entity.Song, error) {
	// The query select
	query := `
        SELECT id, title,album_id,lyrics
        FROM song
        WHERE id = $1`

	var song entity.Song

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific song and then set the result to song variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&song.ID,
		&song.Title,
		&song.Album_id,
		&song.Lyrics,
	)

	// If any error
	if err != nil {
		return nil, err
	}

	return &song, nil
}

// GetAllSong is function to get all songs from database
func (repo *songConnection) GetAllSong(ctx context.Context, offset string, limit string, filter_album_id string) ([]entity.Song, error) {
	// The query select
	var query string
	var rows *sql.Rows
	var err error
	fmt.Println(filter_album_id)
	if filter_album_id == "nothing" {
		query = "SELECT id, title,album_id,lyrics FROM song ORDER BY id LIMIT $2 OFFSET $1"
	} else {
		query = "SELECT id, title,album_id,lyrics FROM song WHERE album_id=$3 ORDER BY id LIMIT $2 OFFSET $1"
	}

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	// Run the query to get all songs
	if filter_album_id == "nothing" {
		rows, err = repo.db.QueryContext(ctx, query, offset, limit)
	} else {
		rows, err = repo.db.QueryContext(ctx, query, offset, limit, filter_album_id)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the songs is not empty
	for rows.Next() {
		var song entity.Song

		// Set to the song variable
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Album_id,
			&song.Lyrics,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the song to songs variable
		songs = append(songs, song)
	}

	return songs, nil
}

// BatchCreate is function to insert some songs in once to database
func (repo *songConnection) BatchCreate(ctx context.Context, songs []entity.Song) ([]int64, error) {
	var IDs []int64

	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return IDs, nil
	}
	// If any error, the transaction will be rollback
	defer tx.Rollback()

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query insert
	query := `INSERT INTO song (title,album_id,lyrics) VALUES ($1) RETURNING id`

	// Loop every song
	for _, song := range songs {
		var id int64

		// Run query insert of every song to database
		err := tx.QueryRowContext(ctx, query, song.Title, song.Album_id, song.Lyrics).Scan(&id)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}

		// Add the new id to IDs variable
		IDs = append(IDs, id)
	}

	// Commit the transaction
	err = tx.Commit()
	// If any error
	if err != nil {
		return IDs, err
	}

	return IDs, nil
}

// Update is function to update song in database
func (repo *songConnection) Update(ctx context.Context, song entity.Song) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE song set title=$1, lyrics=$2 WHERE id=$3`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, song.Title, song.Lyrics, song.ID)
	if err != nil {
		return err
	}

	// Get how many data has been updated
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

// Delete is function to delete song in database
func (repo *songConnection) Delete(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from song WHERE id=$1`

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}
