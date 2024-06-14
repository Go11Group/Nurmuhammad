package models

import (
	"time"
)

type User struct {
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt,omitempty"`
}

type UserGetAllResp struct {
	Users []User
	Count int
}

type UserCourses struct {
	UserId  string    `json:"user_id"`
	Courses []Course2 `json:"courses"`
}

type User2 struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type SearchUser struct {
	Results []User2 `json:"results"`
}
