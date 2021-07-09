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
	taskString, err := arrayToString(ctx, daily.Tasks)
	if err != nil {
		return err
	}

	// Create
	p.db.Create(&repository.Daily{
		Email: string(daily.Email),
		Date:  daily.Date,
		Tasks: taskString,
	})
	return ctx.JSON(http.StatusCreated, daily)
}

func (p *Daily) GetDaily(ctx echo.Context, id int) error {
	// get data
	m := new(repository.Daily)
	if tx := p.db.First(m, "id = ?", id); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}

	//jsonを構造体に変換
	taskArray, err := stringToArray(ctx, m.Tasks)
	if err != nil {
		return err
	}

	daily := &gen.Daily{
		Email: types.Email(m.Email),
		Date:  m.Date,
		Tasks: taskArray,
	}

	return ctx.JSON(http.StatusOK, daily)
}

func arrayToString(ctx echo.Context, array *[]gen.Task) (string, error) {
	b, err := json.Marshal(array)
	if err != nil {
		return "", sendError(ctx, http.StatusBadRequest, "Invalid format")
	}
	return string(b), nil
}

func stringToArray(ctx echo.Context, str string) (*[]gen.Task, error) {
	b := []byte(str)
	sl := make([]gen.Task, 0)
	err := json.Unmarshal(b, &sl)
	if err != nil {
		return nil, sendError(ctx, http.StatusBadRequest, "Invalid format")
	}
	return &sl, nil
}
