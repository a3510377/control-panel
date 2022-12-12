package main

import (
	"fmt"

	"github.com/a3510377/control-panel/database"
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
	// data := &models.Instance{ID: 3864821731328, Name: "\"; DROP TABLE instance_tags; --"}
	fmt.Println(db.GetInstanceByTags("a"))
	// db.Model(&models.Instance{}).Association("Tags").Append(&models.InstanceTags{Name: "awa"})
	// fmt.Println(data)
}
