package database

import (
	"errors"

	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"gorm.io/gorm"
)

func (db *DB) GetInstanceByID(id id.ID) *models.Instance {
	instance := &models.Instance{}

	if err := db.Where("id = ?", id).First(instance).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return instance
}

func (db *DB) GetInstanceTagsByName(name ...string) []int {
	tags, instanceTags := []int{}, []models.InstanceTags{}

	db.Where("name IN ?", name).Find(&instanceTags)
	for _, tag := range instanceTags {
		tags = append(tags, tag.ID)
	}

	return tags
}

func (db *DB) GetInstanceByTags(tags ...string) []models.Instance {
	instances, instanceIds := []models.Instance{}, db.GetInstanceTagsByName(tags...)

	db.Where("id IN ?", instanceIds).Find(&instances)

	return instances
}
