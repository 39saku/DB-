package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func GetListReportHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.GetListReport(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func GetReportHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.GetReport(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func CreateReportHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.CreateReport(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func UpdateReportHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.UpdateReport(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
func DeleteReportHandler(db *sql.DB, c *gin.Context) {
	str, num, err := usecase.DeleteReport(db, c)
	if err != nil {
		c.JSON(num, err)
	} else {
		c.JSON(num, str)
	}
}
