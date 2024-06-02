package main

import (
	"fmt"
	"new/cruds"
	"new/structs"
)

func main() {
	// databsega ulash
	db, err := cruds.ConnectDb()
	if err != nil {
		panic(fmt.Errorf("error to connect db"))
	}
	defer db.Close()

	// har bir table uchun alohida db

	user := cruds.ConnectUserRepo(db)
	product := cruds.ConnectProductRepo(db)
	user_product := cruds.ConnectUserProductRepo(db)

	// product structiga ma'lumotlar yaratib slicega solish va insert qilish
	products := []structs.Product{}
	product1 := structs.Product{Name: "Apple", Description: "red and tasty", Price: 23000.00, Stock_quantity: 0}
	product2 := structs.Product{Name: "Banana", Description: "yellow and sweet", Price: 12000.00, Stock_quantity: 50}
	product3 := structs.Product{Name: "Orange", Description: "juicy and sour", Price: 15000.00, Stock_quantity: 100}
	product4 := structs.Product{Name: "Mango", Description: "tropical and juicy", Price: 20000.00, Stock_quantity: 75}
	product5 := structs.Product{Name: "Pineapple", Description: "tangy and sweet", Price: 18000.00, Stock_quantity: 20}
	products = append(products, product1, product2, product3, product4, product5)
	for _, v := range products {
		err = product.CreateProduct(v)
		if err != nil {
			panic(fmt.Errorf("error on insert product"))
		}
	}
	fmt.Println("Products created succesfully")

	// user structiga ma'lumotlar yaratib slicega solish va insert qilish

	users := []structs.User{
		{Name: "Jack", Email: "Jack@gmail.com", Password: "any"},
		{Name: "Jill", Email: "Jill@gmail.com", Password: "any"},
		{Name: "John", Email: "John@gmail.com", Password: "any"},
		{Name: "Jane", Email: "Jane@gmail.com", Password: "any"},
		{Name: "Bob", Email: "Bob@gmail.com", Password: "any"},
		{Name: "Alice", Email: "Alice@gmail.com", Password: "any"},
		{Name: "Tom", Email: "Tom@gmail.com", Password: "any"},
		{Name: "Jerry", Email: "Jerry@gmail.com", Password: "any"},
		{Name: "Sam", Email: "Sam@gmail.com", Password: "any"},
		{Name: "Sally", Email: "Sally@gmail.com", Password: "any"},
	}
	for _, v := range users {
		err = user.CreateUser(v)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Users created succesfully")

	// user_products tablega ma'lumot qo'shish uchun misol

	// usersproducts := []structs.User_Product{
	// 	{User_id: 1, Product_id: 1},
	// 	{User_id: 1, Product_id: 2},
	// }
	// for _, v := range usersproducts {
	// 	user_product.Create(v)
	// }

	// user va productni user_products table orqali join qilib o'qib olish
	usersandproducts, err := user_product.GetUserProduct()
	if err != nil {
		fmt.Println("Error on read user_products table")
	}
	for _, v := range usersandproducts {
		fmt.Println(v)
	}

}
