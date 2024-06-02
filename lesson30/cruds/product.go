package cruds

import (
	"database/sql"
	"fmt"
	"new/structs"
)

type ProductRepo struct {
	Db *sql.DB
}

func ConnectProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (us *ProductRepo) CreateProduct(product structs.Product) error {
	_, err := us.Db.Exec(`insert into products(name,description,price,stock_quantity)
	values ($1,$2,$3,$4)`, product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil {
		return err
	}
	return nil
}

func (us *ProductRepo) GetProduct() ([]structs.Product, error) {
	products := []structs.Product{}
	rows, err := us.Db.Query(`select * from products`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := structs.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	fmt.Println("Product read succesfully")
	return products, nil
}

func (us *ProductRepo) UpdateProduct(product structs.Product) error {
	_, err := us.Db.Exec(`update products set name=$1,description=$2,price=$3,stock_quantity=$4
	where id=$5`, product.Name, product.Description, product.Price, product.Stock_quantity, product.Id)
	if err != nil {
		return err
	}
	fmt.Println("Product updated succesfully")
	return nil
}

func (us *ProductRepo) DeleteProduct(id int) error {
	_, err := us.Db.Exec(`delete from products where id=$1`, id)
	if err != nil {
		return err
	}
	fmt.Println("Product deleted succesfully")
	return nil
}
