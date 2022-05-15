package user

import "goblog/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty" gorm:"size:100"`
	Email    string `json:"-" gorm:"size:100"`
	Phone    string `json:"-" gorm:"size:20"`
	Password string `json:"-" gorm:"size:100"`

	models.CommonTimestampsField
}
