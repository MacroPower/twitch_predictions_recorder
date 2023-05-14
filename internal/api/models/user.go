package models

type User struct {
	UserID          string `json:"user_id" gorm:"primarykey"`
	UserDisplayName string `json:"user_display_name"`
}
