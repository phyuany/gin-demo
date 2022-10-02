package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp default CURRENT_TIMESTAMP"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp default CURRENT_TIMESTAMP"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uid)
}
