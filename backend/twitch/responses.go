package twitch

type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	ProfileImage    string `json:"profile_image_url"`
	OfflineImage    string `json:"offline_image_url"`
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	BroadcasterType string `json:"broadcaster_type"`
	StartedAt       string `json:"started_at"`
	ViewCount       int    `json:"view_count"`
}

type UsersResponse struct {
	Data []User `json:"data"`
}
type FollowersResponse struct {
	Total int `json:"total"`
}

type StreamInformationResponse struct {
	Data []StreamInformation `json:"data"`
}

type StreamInformation struct {
	// ID              string    `json:"id"`
	UserDisplayName string `json:"user_name"`
	GameName        string `json:"game_name"`
	ViewerCount     int    `json:"viewer_count"`
	Title           string `json:"title"`
	ThumbnailUrl    string `json:"thumbnail_url"`
	StartedAt       string `json:"started_at"`
}

type DetailInformation struct {
	// ID              string    `json:"id"`
	UserDisplayName string `json:"user_name"`
	Description     string `json:"description"`
	GameName        string `json:"game_name"`
	ViewerCount     int    `json:"viewer_count"`
	Title           string `json:"title"`
	ThumbnailUrl    string `json:"thumbnail_url"`
	StartedAt       string `json:"started_at"`
}
