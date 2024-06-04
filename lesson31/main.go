package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-faker/faker/v3"
	_ "github.com/lib/pq"
)

type z struct{}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:pass@localhost:5432/new?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// err = Insert(db)
	// if err != nil {
	// 	panic(fmt.Errorf("error on insert values"))
	// }

	Testing(db)
}

func Insert(db *sql.DB) error {
	for i := 0; i <= 1000000; i++ {
		_, err := db.Exec(`insert into users(username,usersurname,email,password,phonenumber) values ($1,$2,$3,$4,$5)`,
			faker.FirstName(), faker.LastName(), faker.Email(), faker.Password(), faker.Phonenumber())
		if err != nil {
			return err
		}
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func Testing(db *sql.DB) {
	var a int
	v := map[int]z{}
	q := z{}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	for i := 0; i <= 500; i++ {
		go func() {
			// year := faker.FirstName()
			phone := faker.Phonenumber()
			// sur := faker.LastName()
			t := time.Now()
			err := db.QueryRow("select count(1) from users where  phonenumber=$1",
				phone).Scan(&a)
			if err != nil {
				fmt.Println(err)
				return
			}
			v[i] = q
			fmt.Println(i, a, time.Now().Sub(t))
		}()
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
	time.Sleep(3 * time.Second)
	fmt.Println("rtyui", len(v))
}
