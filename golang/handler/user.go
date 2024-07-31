package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(db *sql.DB, c *gin.Context) {
	usecase.CreateUser(db, c)
}
func GetUserHangler(db *sql.DB, c *gin.Context) {
	usecase.GetUser(db, c)
}
func GetUserIdHandler(db *sql.DB, c *gin.Context) {
	usecase.GetUserId(db, c)
}
func UpdateUserHandler(db *sql.DB, c *gin.Context) {
	usecase.UpdateUser(db, c)
}
