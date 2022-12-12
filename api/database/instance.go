package database

import (
	"errors"

	"github.com/a3510377/control-panel/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *DB) GetInstance(data *models.Instance) (*models.Instance, error) {
	result, error := db.FirstInstance(data)
	if error != nil {
		return nil, error
	}

	return result, nil
}

func (db *DB) FindInstance(data *models.Instance) (*models.Instance, error) {
	result := db.Preload(clause.Associations).Find(data)

	return data, result.Error
}

func (db *DB) FirstInstance(data *models.Instance) (*models.Instance, error) {
	err := db.Preload(clause.Associations).First(data, data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return data, err
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
