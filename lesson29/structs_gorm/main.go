package main

import (
	"fmt"
	"new/querys"
	"new/structs"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// connecting to db

	db := querys.ConnectDb()
	us := querys.CreateDb(db)

	// create table

	us.CreateTable(&structs.User{})

	// insert to table

	user := structs.User{FirstName: "Behzod", LastName: "Avazbekovna", Gender: "Male", Email: "Qo'chqorovBekzod@gmail.com", IsEmployee: true, Age: 18}
	user2 := structs.User{FirstName: "Sanjarbek", LastName: "Abdurahmanov", Gender: "Male", Email: "Sanjar@gmail.com", IsEmployee: true, Age: 16}
	user3 := structs.User{FirstName: "Diyorbek", LastName: "Yaxshi bola", Gender: "Male", Email: "Diyor@gmail.com", IsEmployee: false, Age: 18}

	var users []structs.User
	users = append(users, user)
	users = append(users, user2)
	users = append(users, user3)

	for _, v := range users {
		us.InsertToTable(&v)
	}

	// delete all columns
	// us.DeleteAllTable(&structs.User{})

	// update qilish
	user5 := structs.User{FirstName: "Faxridin", LastName: "Rahimboyev", Gender: "Male", Email: "Qo'chqorovBekzod@gmail.com", IsEmployee: true, Age: 19}
	us.UpdateById(1, &user5, &structs.User{})

	// read user by id
	var user6 structs.User
	us.TakeFirstUser(2, &user6)
	fmt.Println(user6)
}
