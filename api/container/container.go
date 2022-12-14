package container

import (
	"github.com/a3510377/control-panel/database"
)

type Container struct {
	*database.DB
}

func NewContainer(db *database.DB) *Container {
	return &Container{db}
}
