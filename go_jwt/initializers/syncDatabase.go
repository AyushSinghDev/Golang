package initializers

import "go_jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
