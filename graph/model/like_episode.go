package model

import "time"

//LikeEpisodeの型
type LikeEpisode struct {
	ID           int64       `json:"id"`
	CreatedAt    time.Time   `json:"created_at"`
	UserID       int64       `json:"user_id"`
	EpisodeID    int64       `json:"episode_id"`
}

//LikeEpisodeに値を受け取るための型
type NewLikeEpisode struct {
	UserID       int64       `json:"user_id"`
	EpisodeID    int64       `json:"episode_id"`
}