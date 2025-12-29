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
	UpdatedAt     time.Time `json:"updated_at"`
}

type StravaActivityResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	Distance         float64  `json:"distance"`
	MovingTime       int      `json:"moving_time"`
	ElapsedTime      int      `json:"elapsed_time"`
	Type             *string  `json:"type"`
	SportType        *string  `json:"sport_type"`
	LocationCity     *string  `json:"location_city"`
	LocationCountry  *string  `json:"location_country"`
	LocationState    *string  `json:"location_state"`
	AverageSpeed     *float64 `json:"average_speed"`
	AverageCadence   *float64 `json:"average_cadence"`
	AverageHeartrate *float64     `json:"average_heartrate"`
	Kilojoules       *float64 `json:"kilojoules"`

	MaxSpeed     *float64 `json:"max_speed"`
	MaxHeartrate *int     `json:"max_heartrate"`

	StartDate      time.Time `json:"start_date"`
	StartDateLocal time.Time `json:"start_date_local"`
}
