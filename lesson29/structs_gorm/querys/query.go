package querys

import (
	"fmt"

	"gorm.io/gorm"
)

type Data struct {
	Db *gorm.DB
}

func CreateDb(db *gorm.DB) Data {
	return Data{Db: db}
}

func (u *Data) CreateTable(info interface{}) {
	u.Db.AutoMigrate(info)
	fmt.Println("Succes to create Table")
}

func (u *Data) InsertToTable(info interface{}) {
	result := u.Db.Create(info)
	if result.Error != nil {
		fmt.Printf("Failed to insert record: %v\n", result.Error)
		return
	}
	fmt.Println("Succes to insert")
}

func (u *Data) DeleteAllTable(info interface{}) {
	result := u.Db.Where("id%1=0").Delete(info)
	if result.Error != nil {
		fmt.Printf("Failed to delete record: %v\n", result.Error)
		return
	}
	fmt.Println("Succes to delete all")
}

func (u *Data) DeleteById(id string, info interface{}) {
	result := u.Db.Where("id=?", id).Delete(info)
	if result.Error != nil {
		fmt.Printf("Failed to delete record: %v\n", result.Error)
		return
	}
	fmt.Println("Succes to delete")
}

func (u *Data) UpdateById(id int, info interface{}, tip interface{}) {
	result := u.Db.Model(tip).Where("id=?", id).Updates(info)
	if result.Error != nil {
		fmt.Printf("Failed to delete record: %v\n", result.Error)
		return
	}
	fmt.Println("Succes to update")
}

func (u *Data) TakeFirstUser(id int, tip interface{}) {
	result := u.Db.First(tip, id)
	if result.Error != nil {
		fmt.Printf("Failed to delete record: %v\n", result.Error)
		return
	}
	fmt.Println("Succes to read user")
}
