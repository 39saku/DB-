package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func GetListReportHandler(db *sql.DB, c *gin.Context) {
	reports, _ := usecase.GetListReport(db, c)

	c.JSON(200, reports)
}
func GetReportHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.GetReport(db, c)
	c.JSON(200, str)
}
func CreateReportHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.CreateReport(db, c)
	c.JSON(200, str)
}
func UpdateReportHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.UpdateReport(db, c)
	c.JSON(200, str)
}
func DeleteReportHandler(db *sql.DB, c *gin.Context) {
	str, _ := usecase.DeleteReport(db, c)
	c.JSON(200, str)
}
