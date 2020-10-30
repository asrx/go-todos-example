package models

import "Todos/dao"

// Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

/**
	Todo Model 的增删改查操作
 */

func CreateTodo(todo *Todo) (err error)  {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (list []*Todo, err error) {
	err = dao.DB.Find(&list).Error
	return
}

func GetATodo(id string) (todo *Todo, err error)  {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id string) (err error)  {
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}