package response

import "time"

type UserMealResponse struct {
	ID       string `json:"id"`
	FoodName string `json:"food_names"`
	Calories float64
	Protein  float64
	Fat      float64
	Carbs    float64

	UserID string        `json:"user_id"`
	User   *UserResponse `json:"user"`

	CreatedAt time.Time `json:"created_at"`
}
