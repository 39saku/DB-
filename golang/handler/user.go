package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.CreateUser(db, c)
	c.JSON(200, str)
}
func GetUserHangler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.GetUser(db)
	c.JSON(200, str)
}
func GetUserIdHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.GetUserId(db, c)
	c.JSON(200, str)
}
func UpdateUserHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.UpdateUser(db, c)
	c.JSON(200, str)
}
