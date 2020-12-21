package models

import (
	"encoding/gob"
	"time"
)

func init() {
	//序列
	gob.Register(User{})
}

// 用户结构体 数据表结构
type User struct {
	Uuid          string 		  `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Status      int           `json:"status"`
	Avatar      string        `json:"avatar"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

//type Password struct {
//	Password          string `json:"password" binding:"required,min=1,max=16"`
//	RePassword        string `json:"re_password" binding:"required,min=8,max=16"`
//	ReConfirmPassword string `json:"re_confirm_password" binding:"eqfield=RePassword"`
//}
//
//type Update struct {
//	Name        string    `json:"name" binding:"min=4,max=12"`
//	Description string    `json:"description" binding:"min=1,max=255"`
//	Status      int       `json:"status"`
//	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updatedAt"`
//}
//
//type Active struct {
//	Status      int       `json:"status"`
//	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updatedAt"`
//}

//type ActiveUrl struct {
//	Url  string
//	Info string
//}
//
//type ActiveInfo struct {
//	Id     string
//	Expire int64
//}