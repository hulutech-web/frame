package migration

import (
	"github.com/hulutech-web/frame/cmd"
	"github.com/hulutech-web/frame/database/migration"
)

func init() {
	cmd.Add(&MigrationInit{})
}

type MigrationInit struct {
}

func (mi *MigrationInit) Command() string {
	return "migration:init"
}

func (mi *MigrationInit) Description() string {
	return "complete a task on the list"
}

func (mi *MigrationInit) Handler(arg *cmd.Arg) error {
	m := &migration.MigrationUtils{}
	m.SetUp()
	return nil
}
