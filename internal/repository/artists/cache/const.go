package cache

import "time"

const (
	artistsKey      = "artists"
	artistDetailKey = "artist:%d"
	expiration      = time.Hour * 1
)
