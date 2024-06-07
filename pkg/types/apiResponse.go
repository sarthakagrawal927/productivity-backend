package types

type DayLevelFoodConsumption struct {
	Name       string `json:"name"`
	Kcal       uint   `json:"kcal"`
	Protein    uint   `json:"protein"`
	Fiber      uint   `json:"fiber"`
	Quantity   uint   `json:"quantity"`
	Created_at string `json:"CreatedAt"`
}
