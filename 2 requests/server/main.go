package main

import (
	"exam/dbcon"
	"exam/gin"
	"fmt"
)

func main() {

	// there i am connect to postgres database
	db, err := dbcon.ConnnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	// and closing database
	defer db.Close()

	// on there i am connecting each table to their repo
	user := dbcon.ConnectUser(db)
	lesson := dbcon.ConnectLesson(db)
	course := dbcon.ConnectCourse(db)
	enrollment := dbcon.ConnectEnrollment(db)

	// there i am connecting to gin and run it
	c := gin.ConnectGin(&gin.Handler{User: user, Lesson: lesson, Course: course, Enrollment: enrollment})
	c.Run()
}
