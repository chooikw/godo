package todoservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func FindMany(userId string) (todos []Todo) {
	db.Where("user_id =?", userId).Find(&todos)
	return
}

func Create(todo Todo) (Todo, error) {
	res := db.Create(&todo)
	return todo, res.Error
}

func Update(id int, data UpdateInput, userId string) (Todo, error) {
	// Allow user to update completed field only
	res := db.Model(&Todo{}).Where("id = ?", id).Where("user_id =?", userId).Updates(map[string]interface{}{"completed": data.Completed, "updatedAt": time.Now()})
	if res.Error != nil {
		return Todo{}, res.Error
	}

	if res.RowsAffected == 0 {
		return Todo{}, errors.New("Record not found")
	}

	todo, findErr := FindOne(id)

	return todo, findErr
}

func FindOne(id int) (Todo, error) {
	var todo Todo
	res := db.Where("id = ?", id).First(&todo)
	return todo, res.Error
}

func Delete(id int, userId string) error {
	var todo Todo
	res := db.Where("id = ?", id).Where("user_id =?", userId).Delete(&todo)
	if res.RowsAffected == 0 {
		return errors.New("Record not found")
	}
	return res.Error
}

func Init(d *gorm.DB) {
	db = d
	db.AutoMigrate(&Todo{})
	fmt.Println("Initialized Todo service")
}
