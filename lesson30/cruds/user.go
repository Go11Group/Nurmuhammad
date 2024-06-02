package cruds

import (
	"database/sql"
	"fmt"
	"new/structs"
)

type UserRepo struct {
	Db *sql.DB
}

func ConnectUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (us *UserRepo) CreateUser(user structs.User) error {
	_, err := us.Db.Exec(`insert into users(username,email,password)
	values ($1,$2,$3)`, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserRepo) GetUser() ([]structs.User, error) {
	users := []structs.User{}
	rows, err := us.Db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := structs.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println("user read succesfully")
	return users, nil
}

func (us *UserRepo) UpdateUser(User structs.User) error {
	_, err := us.Db.Exec(`update users set username=$1,email=$2,password=$3
	where id=$4`, User.Name, User.Email, User.Password, User.Id)
	if err != nil {
		return err
	}
	fmt.Println("user updated succesfully")
	return nil
}

func (us *UserRepo) DeleteUser(id int) error {
	_, err := us.Db.Exec(`delete from users where id=$1`, id)
	if err != nil {
		return err
	}
	fmt.Println("user deleted succesfully")
	return nil
}
