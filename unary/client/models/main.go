package models

type Work struct {
	Name     string `json:"name" validate:"required"`
	Priority int    `json:"priority" validate:"required"`
}
