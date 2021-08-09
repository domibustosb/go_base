package models

import "github.com/domibustosb/go_base/pkg/datalayers"

type Anime struct {
	ID   uint    `gorm:"primary_key"`
	Name *string `gorm:"primary_key"`
}

type AnimeModel struct {
}

func (b AnimeModel) FindAll(ID *uint) (*[]Anime, error) {
	var anime []Anime
	db := datalayers.GetDB()
	result := db.Find(&anime)
	return &anime, result.Error
}
