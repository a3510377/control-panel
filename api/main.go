package main

import (
	"fmt"

	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/models"
)

func main() {
	db, _ := database.NewDB("test.db")
	// i := &models.Instance{}
	// db.Create(i)
	// instance.AddTag("test1", "test2")

	// Instances

	// data := []models.Instance{}
	// .Preload(clause.Associations).Find(&data)
	// fmt.Println(db.Preload(clause.Associations).Find(&data).Error)
	// fmt.Println(data)

	data2 := []models.Tag{}
	fmt.Println(db.Model(&models.Tag{}).Where("name = ?", "test1").Find(&data2).Error)
	for _, i := range data2 {
		fmt.Println(*i.Instances[0])
	}

	// if err := cli.NewCommand().Execute(); err != nil {
	// 	os.Exit(1)
	// }
}

// SELECT count(*) FROM `instances` JOIN `instance_tags` ON `instance_tags`.`instance_id` = `instances`.`id` AND `instance_tags`.`tag_id` IN (NULL)
