package database

import (
	"errors"
	"fmt"
	"time"

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

func (db *DB) getInstanceByID(id id.ID) *models.Instance {
	instance := &models.Instance{}

	if err := db.Where("id = ?", id).First(instance).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return instance
}

func (db *DB) GetInstanceByID(id id.ID) *DBInstance {
	if data := db.getInstanceByID(id); data != nil {
		return &DBInstance{db, *data}
	}
	return nil
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

func (i *DBInstance) GetNow()       { i.Instance = *i.Db.getInstanceByID(i.ID) }
func (i *DBInstance) Save() error   { return i.Db.Save(&i.Instance).Error }
func (i *DBInstance) Get() *gorm.DB { return i.Db.Model(&models.Instance{}).Where("id = ?", i.ID) }

func (i *DBInstance) SetNull(key string) error { return i.Get().Update(key, gorm.Expr("NULL")).Error }

func (i *DBInstance) ClearEndAt() error { return i.SetNull("end_at") }

func (i *DBInstance) SetName(name string) {
	i.Name = name
	i.Save()
}

func (i *DBInstance) SetEndAt(time time.Time) {
	i.EndAt = time
	i.Save()
}
