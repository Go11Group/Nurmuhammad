package models

import "time"

type Lesson struct {
	LessonID  string    `json:"lessonId"`
	CourseID  string    `json:"courseId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type LessonGetAllResp struct {
	Lesson []Lesson
	Count  int
}

type Lesson2 struct {
	LessonID string `json:"lessonId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
