package http

import (
	"DailyreportApi/internal/http/gen"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type Api struct {
}

func (a Api) AddDaily(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) GetDaily(ctx echo.Context, id int) error {
	panic("implement me")
}

func NewApi(db *gorm.DB) *Api {
	return &Api{}
}

var _ gen.ServerInterface = (*Api)(nil)
