package user

import (
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(password string) bool {

	return hash.BcryptCheck(password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
