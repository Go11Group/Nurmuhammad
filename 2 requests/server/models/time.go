package models

type TimePeriod struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type PopularCourse struct {
	CourseID         string `json:"course_id"`
	CourseTitle      string `json:"course_title"`
	EnrollmentsCount int    `json:"enrollments_count"`
}

type ResponseData struct {
	TimePeriod     TimePeriod      `json:"time_period"`
	PopularCourses []PopularCourse `json:"popular_courses"`
}
