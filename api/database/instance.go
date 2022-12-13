package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	if err := db.Where("id = ?", id).Preload(clause.Associations).First(
		instance).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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
	tags, instanceTags := []int{}, []models.Tags{}

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

/* DBInstance */
func (i *DBInstance) GetNow()       { i.Instance = *i.Db.getInstanceByID(i.ID) }
func (i *DBInstance) Get() *gorm.DB { return i.Db.Model(&models.Instance{ID: i.ID}) }

func (i *DBInstance) SetNull(key string) error { return i.Get().Update(key, gorm.Expr("NULL")).Error }

// Tags []Tags `gorm:"many2many:instanceTags;foreignKey:ID;References:ID"` // 標籤

func (i *DBInstance) SetName(name string) error       { return i.Get().Update("name", name).Error }
func (i *DBInstance) SetRootDir(root string) error    { return i.Get().Update("RootDir", root).Error }
func (i *DBInstance) SetType(Type string) error       { return i.Get().Update("RootDir", Type).Error }
func (i *DBInstance) SetLastTime(cmd string) error    { return i.Get().Update("LastTime", cmd).Error }
func (i *DBInstance) SetEndAt(time time.Time) error   { return i.Get().Update("EndAt", time).Error }
func (i *DBInstance) ClearEndAt(time time.Time) error { return i.SetNull("EndAt") }

func (i *DBInstance) SetStartCommand(cmd string) error {
	return i.Get().Update("StartCommand", cmd).Error
}

func (i *DBInstance) SetStopCommand(cmd string) error {
	return i.Get().Update("StopCommand", cmd).Error
}

func (i *DBInstance) GetTags() []string {
	tags, strTags := []models.Tags{}, []string{}

	i.Get().Association("Tags").Find(&tags)
	for _, tag := range tags {
		strTags = append(strTags, tag.Name)
	}

	return strTags
}

func (i *DBInstance) AddTag(tags ...string) {
	tagsList := []models.Tags{}
	oldTags := i.GetTags()

	for _, tag := range tags {
		if slices.Contains(oldTags, tag) || tag == "" {
			continue
		}
		tagsList = append(tagsList, models.Tags{Name: tag})
	}

	i.Get().Association("Tags").Append(tagsList)
}

func (i *DBInstance) RemoveTag(tag string) {
	i.Get().Association("Tags").Delete(&models.Tags{Name: tag})
}
