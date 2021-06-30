package repository

type DailyData struct {
	ID    uint64 `gorm:"primaryKey"`
	Email string
	Date  string
	Tasks string
}
