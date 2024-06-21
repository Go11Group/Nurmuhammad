package models

import "time"

type Course struct {
	CourseID    string    `json:"courseId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

type Course2 struct {
	CourseID    string `json:"courseId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseGetAllResp struct {
	Course []Course
	Count  int
}

type CourseLessons struct {
	CourseId string
	Lessons  []Lesson2
}

type CourseUsers struct {
	CourseId string
	Users    []User2
}
