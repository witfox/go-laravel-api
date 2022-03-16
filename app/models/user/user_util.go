package user

import "gohub/pkg/database"

func GetByMulti(account string) (userModel User) {
	database.DB.Where("phone = ?", account).
		Or("email = ?", account).Or("name = ?", account).First(&userModel)
	return
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

func GetByEamil(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

func IsEmailExists(email string) bool {

	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)

	return count > 0
}

func IsPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}

func Get(idStr string) (userModel User) {
	database.DB.Where("id=?", idStr).First(&userModel)
	return
}
