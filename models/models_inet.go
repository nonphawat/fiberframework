package models

import "gorm.io/gorm"

type User struct {
	Email        string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Username     string `json:"username" validate: "required"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	Phone_number string `json:"ph_num" validate:"required"`
	Business     string `json:"business" validate:"required"`
	Web_name     string `json:"web_name" validate:"required,min=3,max=30"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name"`
	Fund     int    `json:"fund"`
	Employee int    `json:"employee"`
	Email    string `json:"email"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data  []DogsRes `json:"data"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
}

type ResultDataV2 struct {
	Data    []DogsRes `json:"data"`
	Name    string    `json:"name"`
	Count   int       `json:"count"`
	Red     int       `json:"sum_red"`
	Green   int       `json:"sum_green"`
	Pink    int       `json:"sum_pink"`
	Nocolor int       `json:"sum_nocolor"`
}

type Employee struct {
	gorm.Model
	Employee_id string `json:"employee_id"`
	Fund        int    `json:"fund"`
	Employee    int    `json:"employee"`
	Email       string `json:"email"`
}
