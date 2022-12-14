package database

import (
	"encoding/json"
	"strings"

	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/node"
	"github.com/a3510377/control-panel/service/id"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var instancesNode = map[id.ID]*node.Instance{}

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

func (db *DB) CreateNewInstance(instances *models.Instance) (*DBInstance, error) {
	if err := db.Create(instances).Error; err != nil {
		return nil, err
	}
	return &DBInstance{db, *instances}, nil
}

// 獲取全部實例
func (db *DB) GetAllInstances() []DBInstance {
	return ModelInstancesToDBInstances(db, find(db.PreloadAll(), []models.Instance{}))
}

func (db *DB) GetAutoStartInstances() []DBInstance {
	return ModelInstancesToDBInstances(db, find(db.PreloadAll().Where("auto_start = ?", true), []models.Instance{}))
}

// 從實例 ID 獲取實例(s)
func (db *DB) GetInstancesByID(id ...id.ID) []DBInstance {
	return ModelInstancesToDBInstances(db, find(db.PreloadAll(), []models.Instance{}, id))
}

// 從實例 ID 獲取實例
func (db *DB) GetInstanceByID(id id.ID) *DBInstance {
	return &DBInstance{db, find(db.PreloadAll(), models.Instance{}, id)}
}

// 從實例名稱獲取實例
func (db *DB) GetInstanceByName(name string) *DBInstance {
	return &DBInstance{db, find(db.PreloadAll().Where("name = ?", name), models.Instance{})}
}

func (db *DB) GetInstancesByTag(tags ...string) []DBInstance {
	dbTag, instances := find(db.Select("name").Where("name IN ?", tags), models.Tag{}), []DBInstance{}
	for _, instance := range dbTag.Instances {
		instances = append(instances, DBInstance{db, *instance})
	}
	return instances
}

/* DBInstance */

func (i *DBInstance) GetNow() {
	if data := i.Db.GetInstanceByID(i.ID); data != nil {
		i.Instance = data.Instance
	}
}
func (i *DBInstance) Get() *gorm.DB                            { return i.Db.Model(&models.Instance{ID: i.ID}) }
func (i *DBInstance) Model() *models.Instance                  { return &models.Instance{ID: i.ID} }
func (i *DBInstance) Delete() error                            { return i.Db.Delete(i.Model()).Error }
func (i *DBInstance) Update(column string, value any) *gorm.DB { return i.Get().Update(column, value) }
func (i *DBInstance) SetNull(key string) error                 { return i.Update(key, gorm.Expr("NULL")).Error }

func (i *DBInstance) Updates(values any) *gorm.DB {
	return i.Get().Omit("ID").Omit("CreatedAt").Updates(values)
}

func (i *DBInstance) GetTags() []string {
	tags, strTags := []models.Tag{}, []string{}

	i.Get().Association("Tags").Find(&tags)
	for _, tag := range tags {
		strTags = append(strTags, tag.Name)
	}

	return strTags
}

func (i *DBInstance) AddTag(tags ...string) {
	tagsList := []models.Tag{}
	oldTags := i.GetTags()

	for _, tag := range tags {
		if slices.Contains(oldTags, tag) || tag == "" {
			continue
		}
		tagsList = append(tagsList, models.Tag{Name: tag})
	}

	i.Get().Association("Tags").Append(tagsList)
}

func (i *DBInstance) RemoveTag(tag string) {
	i.Get().Association("Tags").Delete(&models.Tag{Name: tag})
}

func (i *DBInstance) Node(tags ...string) *node.Instance {
	data, ok := instancesNode[i.ID]
	if !ok {
		// i.RootDir, i.StartCommand)
		data = &node.Instance{
			ID:          i.ID,
			Name:        i.Name,
			Root:        i.RootDir,
			ProcessArgs: strings.Split(i.StartCommand, " "),
		}
		instancesNode[i.ID] = data
	}
	return data
}

func (i *DBInstance) Run() error { return i.Node().Run() }

func (i *DBInstance) JSON() (result map[string]any) {
	i.GetNow()
	b, _ := json.Marshal(&i.Instance)
	_ = json.Unmarshal(b, &result)
	return
}
