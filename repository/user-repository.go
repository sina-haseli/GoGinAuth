package repository

import (
	"golang.com/m/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  type UserRepository interface {
	GetAllUsers(id int64) entity.User
	CreateUser(user *entity.User) entity.User
  }

  type database struct {
	  connection *gorm.DB
  }

  func NewUserRepository() UserRepository {

	dsn := "host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	  return &database{
		  connection: db,
	  }
  }

func (db *database) GetAllUsers(id int64) entity.User {
	var users entity.User
	db.connection.Find(&users, id)
	return users
}

func (db *database) CreateUser(user *entity.User) entity.User {
	db.connection.Create(&user)
	return *user
}