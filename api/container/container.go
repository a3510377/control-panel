package container

import (
	"github.com/a3510377/control-panel/database"
)

type (
	accountContainer struct {
		db *database.DB
	}
	Container struct {
		*database.DB
		Account accountContainer
	}
)

func NewContainer(db *database.DB) *Container {
	return &Container{db, accountContainer{
		db: db,
	}}
}
