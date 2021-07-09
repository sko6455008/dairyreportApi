package repository

type Daily struct {
	ID    uint64 `gorm:"primaryKey"`
	Email string
	Date  string
	Tasks string
}
