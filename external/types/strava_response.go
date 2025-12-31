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
	Username      *string   `json:"username"`
	Weight        int       `json:"weight"`
	Premium       bool      `json:"premium"`
	Profile       string    `json:"profile"`
	ProfileMedium string    `json:"profile_medium"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type StravaActivityResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	Calories         *float64 `json:"calories"`
	Distance         float64  `json:"distance"`
	MovingTime       int      `json:"moving_time"`
	ElapsedTime      int      `json:"elapsed_time"`
	Type             *string  `json:"type"`
	SportType        *string  `json:"sport_type"`
	AverageHeartrate *float64 `json:"average_heartrate"`
	Kilojoules       *float64 `json:"kilojoules"`
	MaxHeartrate     *int     `json:"max_heartrate"`

	StartDate time.Time `json:"start_date"`
}
