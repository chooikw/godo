package todoservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func FindMany() (todos []Todo) {
	db.Find(&todos)
	return
}

func Create(todo Todo) (Todo, error) {
	res := db.Create(&todo)
	return todo, res.Error
}

func Update(id int, data UpdateInput) (Todo, error) {
	// Allow user to update completed field only
	res := db.Model(&Todo{}).Where("id = ?", id).Updates(map[string]interface{}{"completed": data.Completed, "updatedAt": time.Now()})
	if res.Error != nil {
		return Todo{}, res.Error
	}

	todo, findErr := FindOne(id)

	return todo, findErr
}

func FindOne(id int) (Todo, error) {
	var todo Todo
	res := db.Where("id = ?", id).First(&todo)
	return todo, res.Error
}

func Delete(id int) error {
	var todo Todo
	res := db.Where("id = ?", id).Delete(&todo)
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
