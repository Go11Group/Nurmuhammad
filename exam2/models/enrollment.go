package models

import "time"

type Enrollment struct {
	EnrollmentID   string    `json:"enrollmentId"`
	UserID         string    `json:"userId"`
	CourseID       string    `json:"courseId"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}

type EnrollmentGetAllResp struct {
	Enrollment []Enrollment
	Count      int
}
