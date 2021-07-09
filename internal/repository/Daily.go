package repository

type Daily struct {
	ID    int64 `gorm:"primaryKey"`
	Email string
	Date  string
	Tasks string
}
