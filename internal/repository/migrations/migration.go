package migrations

import (
	"gorm.io/gorm"
)

type MigrateDatabase interface {
	Start() error
	DropTable() error
}

type MigrateModel interface {
	TableName() string
}

type migration struct {
	DB     *gorm.DB
	Models []MigrateModel
}

func NewMigration(DB *gorm.DB, models ...MigrateModel) MigrateDatabase {

	return &migration{
		DB:     DB,
		Models: models,
	}
}

func (migration migration) Start() error {

	// Migrate all given models
	for _, model := range migration.Models {
		if err := migration.DB.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

func (migration migration) DropTable() error {

	for _, model := range migration.Models {
		tableName := model.TableName()
		if err := migration.DB.Migrator().DropTable(tableName); err != nil {
			return err
		}
	}
	return nil
}
