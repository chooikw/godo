package todoservice

import "time"

type Todo struct {
	Id        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" binding:"required"`
	Completed bool      `json:"completed"`
	UserId    string    `json:"userId" gorm:"index"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}

type UpdateInput struct {
	Completed bool `json:"completed"`
}
