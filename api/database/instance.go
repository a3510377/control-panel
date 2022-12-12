package database

import (
	"errors"
	"fmt"

	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"gorm.io/gorm"
)

type DBInstance struct {
	Db *DB
	models.Instance
}

func ModelInstancesToDBInstances(db *DB, instances []models.Instance) []DBInstance {
	dbInstances := []DBInstance{}
	for _, instance := range instances {
		dbInstances = append(dbInstances, DBInstance{db, instance})
	}
	return dbInstances
}

func (db *DB) GetInstanceByID(id id.ID) *DBInstance {
	instance := models.Instance{}

	if err := db.Where("id = ?", id).First(&instance).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &DBInstance{db, instance}
}

func (db *DB) GetInstanceTagsByName(name ...string) []int {
	tags, instanceTags := []int{}, []models.InstanceTags{}

	db.Where("name IN ?", name).Find(&instanceTags)
	for _, tag := range instanceTags {
		tags = append(tags, tag.ID)
	}

	return tags
}

func (db *DB) getInstanceByTags(tags ...string) *gorm.DB {
	return db.Where("id IN ?", db.GetInstanceTagsByName(tags...))
}

func (db *DB) GetInstanceByTags(tags ...string) []DBInstance {
	instances := []models.Instance{}

	db.getInstanceByTags(tags...).Find(&instances)

	return ModelInstancesToDBInstances(db, instances)
}

func (db *DB) GetInstanceByNameAndTags(name string, tags []string) []DBInstance {
	instances := []models.Instance{}

	db.getInstanceByTags(tags...).Where("name LIKE ?", fmt.Sprintf("%v%%", name)).Find(&instances)

	return ModelInstancesToDBInstances(db, instances)
}
