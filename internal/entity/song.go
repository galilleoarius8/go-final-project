package entity

type Song struct {
	ID       int64  `json:"id"`
	Album_id int64  `json:"album_id"`
	Title    string `json:"title"`
	Lyrics   string `json:"lyrics"`
}
