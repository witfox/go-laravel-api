package user

import "gohub/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(60);not null;unique;" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(40);" json:"-"` //敏感数据不返回
	Phone    string `gorm:"type:varchar(20);" json:"-"`
	Password string `gorm:"type:varchar(100);" json:"-"`

	models.CommonTimeAt
}
