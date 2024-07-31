package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.CreateUser(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func GetUserHangler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.GetUser(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func GetUserIdHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.GetUserId(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func UpdateUserHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.UpdateUser(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
