package entity

type Album struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Artist_id int64  `json:"artist_id"`
	Price     int64  `json:"price"`
}
