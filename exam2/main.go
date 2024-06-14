package main

import (
	"exam/dbcon"
	"exam/gin"
	"fmt"
)

func main() {
	db, err := dbcon.ConnnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	user := dbcon.ConnectUser(db)
	lesson := dbcon.ConnectLesson(db)
	course := dbcon.ConnectCourse(db)
	enrollment := dbcon.ConnectEnrollment(db)
	c := gin.ConnectGin(&gin.Handler{User: user, Lesson: lesson, Course: course, Enrollment: enrollment})
	c.Run()
}
