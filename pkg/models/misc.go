package models

import "gorm.io/gorm"

type Food_Item struct {
	Name    string `json:"name"`
	Kcal    int    `json:"kcal"`
	Protein int    `json:"protein"`
	Fiber   int    `json:"fiber"`
	// can add more macro nutrients later
	gorm.Model
}

type FoodConsumed struct {
	Food_ItemID uint   `json:"food_item_id"`
	Quantity    int    `json:"quantity"`
	Date        string `json:"date"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type Exercise struct {
	Name                string `json:"name"`
	Type                string `json:"type"` // cardio, strength, flexibility
	PrimaryTargetMuscle string `json:"primary_target_muscle"`
}
