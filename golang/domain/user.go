package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   string
	Name string
}
