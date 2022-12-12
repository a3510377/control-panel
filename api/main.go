package main

import (
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/models"
)

func main() {
	db, _ := database.NewDB("test.db")

	instance := models.Instance{}
	db.Create(&instance)
	// db.Model(&instance).Where("id = ?", 4217358376960).Association("Tags").Append(&models.Tags{Name: "b"})
	// db.Model(&instance).Where("id = ?", 4217614946304).Association("Tags").Append(&models.Tags{Name: "b"})
}
