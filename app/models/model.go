package models

import (
	"time"

	"github.com/spf13/cast"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimeAt struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

func (bm BaseModel) GetStringID() string {
	return cast.ToString(bm.ID)
}
