package usecase

import (
	"DailyreportApi/internal/http/gen"
	"DailyreportApi/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Daily struct {
	db *gorm.DB
}

func NewDaily(db *gorm.DB) *Daily {
	return &Daily{
		db: db,
	}
}

func (p *Daily) AddDaily(ctx echo.Context) error {
	// get request
	daily := new(gen.Daily)
	err := ctx.Bind(daily)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format")
	}

	// Create

	p.db.Create(&repository.DailyData{
		Id:    *daily.Id,
		Email: string(daily.Email),
		Date:  daily.Date.Format("2006-01-02"),
	})
	return ctx.JSON(http.StatusCreated, daily)
}

func (p *Daily) GetDaily(ctx echo.Context, id int) error {
	// get data
	m := new(repository.DailyData)
	if tx := p.db.First(m, "id = ?", id); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}
	daily := &gen.Daily{
		Id:    &m.Id,
		Email: m.Email,
		Date:  m.Date,
	}
	return ctx.JSON(http.StatusOK, daily)
}
