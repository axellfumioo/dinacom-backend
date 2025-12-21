package types

type AINutritionResponse struct {
	FoodName      string  `json:"food_name"`
	Calories      int     `json:"calories"`
	Protein       float64 `json:"protein"`
	Carbohydrates float64 `json:"carbohydrates"`
	Fat           float64 `json:"fat"`
}
