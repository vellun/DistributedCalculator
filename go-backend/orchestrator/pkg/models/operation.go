package models

type Operation struct {
	Id       int    `json:"id"`
	Duration int    `json:"duration"`
	Name     string `json:"name"`
}
