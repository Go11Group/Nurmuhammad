package model

import "time"

type User struct {
	UserID         int
	FirstName      string
	LastName       string
	Email          string
	EnrollmentDate time.Time
}

type Problem struct {
	ProblemID   int
	ProblemName string
	Description string
}

type SolvedProblems struct {
	FirstName   string
	LastName    string
	Email       string
	ProblemName string
	Description string
}
