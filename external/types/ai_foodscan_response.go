package types

type AIFoodScanResponse struct {
	Response AINutritionResponse `json:"response"`
}

type AINutritionResponse struct {
	FoodName        string     `json:"food_name"`
	FoodType        string     `json:"food_type"`
	IsFastFood      bool       `json:"is_fast_food"`
	Nutrition       Nutritions `json:"nutrition"`
	CaloriesKcal    int        `json:"calories_kcal"`
	Vitamins        []string   `json:"vitamins"`
	HealthScores    int        `json:"health_scores"`
	HealthNote      string     `json:"health_note"`
	ConfidenceScore float64    `json:"confidence_score"`
}

type Nutritions struct {
	ProteinG float64 `json:"protein_g"`
	CarbsG   float64 `json:"carbs_g"`
	FatG     float64 `json:"fat_g"`
}

// {
//   "response": {
//     "food_name": "Nasi goreng ayam dengan kerupuk",
//     "food_type": "traditional",
//     "is_fast_food": false,
//     "calories_kcal": 720,
//     "nutrition": {
//       "carbs_g": 95,
//       "protein_g": 26,
//       "fat_g": 26
//     },
//     "vitamins": [
//       "Vitamin A",
//       "Vitamin B",
//       "Vitamin C",
//       "Iron",
//       "Potassium"
//     ],
//     "health_score": 55,
//     "health_notes": "Karbohidrat dan lemak cenderung tinggi dari nasi dan minyak tumisan; ada protein ayam dan sedikit sayuran (daun bawang/cabai) namun porsi sayur terbatas. Kerupuk menambah kalori dan natrium.",
//     "confidence_score": 0.78
//   }
// }
