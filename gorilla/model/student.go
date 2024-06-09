package model

import "time"

type Student struct {
	StudentID      int
	FirstName      string
	LastName       string
	Email          string
	DateOfBirth    time.Time
	EnrollmentDate time.Time
}

type Lesson struct {
	LessonID    int
	LessonName  string
	Description string
}

type StudentLesson struct {
	FirstName      string
	LastName       string
	Email          string
	LessonName     string
	Description    string
	DateOfBirth    time.Time
	EnrollmentDate time.Time
}
