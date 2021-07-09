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

func (a *Api) AddDaily(ctx echo.Context) error {
	return a.daily.AddDaily(ctx)
}

func (a *Api) FindDailyById(ctx echo.Context, id int64) error {
	return a.daily.FindDailyById(ctx, id)
}

func (a *Api) FindDaily(ctx echo.Context, params gen.FindDailyParams) error {
	return a.daily.FindDaily(ctx, params)
}

func NewApi(db *gorm.DB) *Api {
	return &Api{daily: usecase.NewDaily(db)}
}

var _ gen.ServerInterface = (*Api)(nil)
