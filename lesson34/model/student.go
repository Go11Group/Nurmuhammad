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
