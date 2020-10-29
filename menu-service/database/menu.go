package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Menu struct {
	ID    string `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price int    `json:"name"`
}

func (menu *Menu) Insert(db *gorm.DB) error {
	result := db.Create(menu)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Berhasil insert menu")
	return nil
}

func (menu *Menu) GetAll(db *gorm.DB) ([]Menu, error) {
	var menus []Menu
	result := db.Find(&menus)
	if result.Error != nil {
		return menus, result.Error
	}

	return menus, nil
}
