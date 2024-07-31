package handler

import (
	"database/sql"
	"db_assignment/usecase"

	"github.com/gin-gonic/gin"
)

func GetListReportHandler(db *sql.DB, c *gin.Context) {
	usecase.GetListReport(db, c)
}
func GetReportHandler(db *sql.DB, c *gin.Context) {
	usecase.GetReport(db, c)
}
func CreateReportHandler(db *sql.DB, c *gin.Context) {
	usecase.CreateReport(db, c)
}
func UpdateReportHandler(db *sql.DB, c *gin.Context) {
	usecase.UpdateReport(db, c)
}
func DeleteReportHandler(db *sql.DB, c *gin.Context) {
	usecase.DeleteReport(db, c)
}
