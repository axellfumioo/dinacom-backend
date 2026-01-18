package payload

type AIInsightPayload struct {
	User                  UserAIInsight         `json:"user"`
	DailyNutritionSummary DailyNutritionSummary `json:"daily_nutrition_summary"`
}

type UserAIInsight struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Smoking       bool   `json:"smoking"`
	SleepDuration int    `json:"sleep_duration"`
	SportDuration string `json:"sport_duration"`
}

type DailyNutritionSummary struct {
	CaloriesKcal float64            `json:"calories_kcal"`
	Nutrition    DailyNutrition `json:"nutrition"`
	Vitamins     []string       `json:"vitamins"`
}

type DailyNutrition struct {
	CarbsG   float64 `json:"carbs_g"`
	ProteinG float64 `json:"protein_g"`
	FatG     float64 `json:"fat_g"`
}

// {
//   "user": {
//     "id": "user_123",
//     "name": "Dodi Mulya",
//     "age": 76,
//     "gender": "male",
//     "smoking": "true",
//     "sleep_duration": 5,
//     "sports_duration": 2
//   },

//   "daily_nutrition_summary": {
//     "calories_kcal": 45,
//     "nutrition": {
//       "carbs_g": 55,
//       "protein_g": 900,
//       "fat_g": 40
//     },
//     "vitamins": [
//       "Vitamin A",
//       "Vitamin C",
//       "Iron",
//       "Potassium"
//     ]
//   },

//   "meta": {
//     "timezone": "Asia/Jakarta",
//     "data_completeness": 0.88
//   }
// }

// ini rsponse nya
// {
//   "health_score": 0,
//   "personal_ai_insight": "",
//   "ai_important_notice": "",
//   "confidence_level": 0
// }
