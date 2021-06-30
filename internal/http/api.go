package http

import (
	"DailyreportApi/internal/http/gen"
	"DailyreportApi/internal/http/usecase"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type Api struct {
	daily *usecase.Daily
}

func (a Api) AddDaily(ctx echo.Context) error {
	return a.daily.AddDaily(ctx)
}

func (a Api) GetDaily(ctx echo.Context, id int) error {
	return a.daily.GetDaily(ctx, id)
}

func NewApi(db *gorm.DB) *Api {
	return &Api{daily: usecase.NewDaily(db)}
}

var _ gen.ServerInterface = (*Api)(nil)
