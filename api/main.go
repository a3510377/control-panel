package main

import (
	"fmt"

	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/models"
)

type Test struct {
	Name string `default:"awa"`
}

func main() {
	db, _ := database.NewDB("test.db")

	// test := &models.InstanceTags{Name: "awa"}
	// fmt.Println(db.Create(test))
	// fmt.Println(db.Create(&models.Instance{
	// 	Name: "test",
	// 	Tags: []models.InstanceTags{{Name: "awa"}},
	// }))
	data := &models.Instance{}
	db.Preload("Tags").Find(data)
	fmt.Println(data)
}
