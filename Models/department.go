package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Dep_name string
	Hod_name string
	Dep_id string
}