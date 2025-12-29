package types

import "time"

type StravaTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	Athlete      any    `json:"athlete"`
}

type StravaProfileResponse struct {
	ID            int64     `json:"id"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Premium       bool      `json:"premium"`
	Profile       string    `json:"profile"`
	ProfileMedium string    `json:"profile_medium"`
	ResourceState int       `json:"resource_state"`
	Sex           string    `json:"sex"`
	State         string    `json:"state"`
	Summit        bool      `json:"summit"`
	Username      *string   `json:"username"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
