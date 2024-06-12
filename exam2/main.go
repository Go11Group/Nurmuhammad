package main

import (
	"exam/dbcon"
	"exam/gin"
)

func main() {
	db, err := dbcon.ConnnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	user := dbcon.ConnectUser(db)
	lesson := dbcon.ConnectLesson(db)
	course := dbcon.ConnectCourse(db)
	enrollment := dbcon.ConnectEnrollment(db)
	c := gin.ConnectGin(&gin.Handler{User: user, Lesson: lesson, Course: course, Enrollment: enrollment})
	c.Run()
}
