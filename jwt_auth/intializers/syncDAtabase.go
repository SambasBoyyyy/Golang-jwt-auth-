package intializers

import "go/jwt_auth/models"

func SyncDAtabase() {
	
	GetDB().AutoMigrate(&models.User{})
}