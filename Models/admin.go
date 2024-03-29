package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model  //is a struct that provides common fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	            //   for database models managed by the GORM (Go Object Relational Mapper) package.
	Name string
	Username string
	Password string
}