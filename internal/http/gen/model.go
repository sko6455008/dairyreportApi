// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package gen

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for TaskProject.
const (
	TaskProject TaskProject = "その他"

	TaskProject1 TaskProject = "おしるこ"

	TaskProject2 TaskProject = "東京メトロ"

	TaskProject3 TaskProject = "ふれあいリズムダンス"

	TaskProject4 TaskProject = "相模屋美術店"

	TaskProjectGC TaskProject = "GC"

	TaskProjectItemstore TaskProject = "itemstore"

	TaskProjectJRA TaskProject = "JRA"

	TaskProjectMtg TaskProject = "社内(mtg/調査など)"
)

// Daily defines model for Daily.
type Daily struct {
	Date  openapi_types.Date  `json:"date"`
	Email openapi_types.Email `json:"email"`
	Id    *int                `json:"id,omitempty"`
	Tasks *[]Task             `json:"tasks,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Task defines model for Task.
type Task struct {
	Hour    float32     `json:"hour"`
	Project TaskProject `json:"project"`
}

// TaskProject defines model for Task.Project.
type TaskProject string

// AddDailyJSONBody defines parameters for AddDaily.
type AddDailyJSONBody Daily

// AddDailyJSONRequestBody defines body for AddDaily for application/json ContentType.
type AddDailyJSONRequestBody AddDailyJSONBody

