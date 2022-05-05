package migrations

import (
	models "strateegy/project-service/models/project"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Project{})
}
