package dao

type Dao interface {
	Create(interface{}) error
	GetAll() ([]interface{}, error)
}

type HabitDao interface {
	Dao
}
