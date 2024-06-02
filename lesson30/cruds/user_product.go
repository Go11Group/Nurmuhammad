package cruds

import (
	"database/sql"
	"new/structs"
)

type UserProductRepo struct {
	Db *sql.DB
}

func ConnectUserProductRepo(db *sql.DB) *UserProductRepo {
	return &UserProductRepo{Db: db}
}

func (us *UserProductRepo) Create(user structs.User_Product) error {
	tr, err := us.Db.Begin()
	if err != nil {
		return err
	}
	_, err = us.Db.Exec(`insert into user_products(user_id,product_id)
	values ($1,$2)`, user.User_id, user.Product_id)
	if err != nil {
		return err
	}
	defer tr.Commit()
	return nil
}

func (us *UserProductRepo) GetUserProduct() ([]structs.UserProductByName, error) {
	users := []structs.UserProductByName{}
	rows, err := us.Db.Query(`select up.id,u.username,p.name 
	from user_products as up
	join users as u
	on u.id=up.user_id
	join products as p 
	on p.id=up.product_id`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := structs.UserProductByName{}
		err = rows.Scan(&user.Id, &user.Name_user, &user.Name_products)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (us *UserProductRepo) Update(user structs.User_Product) error {
	tr, err := us.Db.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	_, err = us.Db.Exec(`update user_products set user_id=$1,product_id=$2 where id=$3`, user.User_id, user.Product_id, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserProductRepo) Delete(id int) error {
	tr, err := us.Db.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	_, err = us.Db.Exec(`delete from user_products where id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
