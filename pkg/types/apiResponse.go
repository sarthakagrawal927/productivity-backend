package types

type DayLevelFoodConsumption struct {
	Name       string  `json:"name"`
	Kcal       float32 `json:"kcal"`
	Protein    float32 `json:"protein"`
	Fiber      float32 `json:"fiber"`
	Fat        float32 `json:"fat"`
	Carbs      float32 `json:"carbs"`
	Quantity   float32 `json:"quantity"`
	Created_at string  `json:"CreatedAt"`
}
