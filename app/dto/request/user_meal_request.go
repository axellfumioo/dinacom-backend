package request

import "backend-dinakom/app/constants"

type CreateUserMealRequest struct {
	FoodName string  `json:"food_name"`
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Carbs    float64 `json:"carbohydrate"`
	Portion int                    `json:"portion"`
	Time    constants.UserMealTime `json:"time"` // BREAKFAST, LUNCH, DINNER, SNACK
}
