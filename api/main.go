package main

import (
	"fmt"

	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/service/id"
)

func main() {
	db, _ := database.NewDB("test.db")

	// instance := models.Instance{}
	// db.Create(&instance)
	// db.Save(&instance)
	// ID := instance.ID
	ID := id.ID(4509189898240)
	fmt.Println(ID)

	// db.Model(&instance).Where("id = ?", ID).Association("Tags").Append(&models.Tags{Name: "b"})
	// db.Model(&instance).Where("id = ?", ID).Association("Tags").Append(&models.Tags{Name: "a"})

	if data := db.GetInstanceByID(ID); data != nil {
		fmt.Println(data.GetTags())
		// data.SetName("awa")
	}
	// .GetTags()
}
