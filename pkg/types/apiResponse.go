package types

type DayLevelFoodConsumption struct {
	Name       string  `json:"name"`
	Kcal       float32 `json:"kcal"`
	Protein    float32 `json:"protein"`
	Fiber      float32 `json:"fiber"`
	Quantity   float32 `json:"quantity"`
	Created_at string  `json:"CreatedAt"`
}
