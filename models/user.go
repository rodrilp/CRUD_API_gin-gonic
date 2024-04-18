package models

type User struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Age  int8   `json:"age" gorm:"type:int"`
}
