package usecase

import (
	"DailyreportApi/internal/http/gen"
	"DailyreportApi/internal/repository"
	"encoding/json"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
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

	// Tasksfieldをjson変換
	s, err := json.Marshal(daily.Tasks)
	if err != nil {
		panic(err)
	}

	// Create
	p.db.Create(&repository.DailyData{
		Email: string(daily.Email),
		Date:  daily.Date.Format("2006-01-02"),
		Tasks: string(s),
	})
	return ctx.JSON(http.StatusCreated, daily)
}

func (p *Daily) GetDaily(ctx echo.Context, id int) error {
	// get data
	m := new(repository.DailyData)
	if tx := p.db.First(m, "id = ?", id); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}

	//jsonを構造体に変換
	tasks := make([]gen.Task, 0)
	b := make([]byte, 0)
	json.Unmarshal(b, &tasks)
	daily := &gen.Daily{
		Email: types.Email(m.Email),
		Date:  types.Date{},
		Tasks: &tasks,
	}

	return ctx.JSON(http.StatusOK, daily)
}
