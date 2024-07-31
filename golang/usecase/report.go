package usecase

import (
	"database/sql"
	"db_assignment/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetListReport(db *sql.DB, context *gin.Context) ([]domain.Report, error) {
	userID := context.Query("id")
	rows, err := db.Query("SELECT * FROM Report WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー: %v", err)
	}
	defer rows.Close()
	var reports []domain.Report
	for rows.Next() {
		var report domain.Report
		err := rows.Scan(&report.Id, &report.User_id, &report.Title, &report.Character_counts, &report.Style, &report.Language)
		if err != nil {
			return nil, fmt.Errorf("行のスキャンエラー: %v", err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("行の反復中のエラー: %v", err)
	}

	return reports, nil
}

func GetReport(db *sql.DB, context *gin.Context) ([]domain.Report, error) {

	Id := context.Query("id")
	rows, err := db.Query("SELECT * FROM Report WHERE id = ?", Id)
	if err != nil {
		return nil, fmt.Errorf("can not Query: %v", err)
	}
	defer rows.Close()

	var reports []domain.Report
	for rows.Next() {
		var report domain.Report
		err := rows.Scan(&report.Id, &report.User_id, &report.Title, &report.Character_counts, &report.Style, &report.Language)
		if err != nil {
			return nil, fmt.Errorf("行のスキャンエラー: %v", err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("行の反復中のエラー: %v", err)
	}

	return reports, nil
}

func CreateReport(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("INSERT INTO Report (user_id, title, character_counts, style, language) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}
	defer state.Close()

	user_id := context.Query("user_id")
	title := context.Query("title")
	characterCounts := context.Query("character_counts")
	style := context.Query("style")
	language := context.Query("language")
	_, err = state.Exec(user_id, title, characterCounts, style, language)
	if err != nil {
		return "", fmt.Errorf("can not Exec: %v", err)
	}
	return "report created", nil
}

func UpdateReport(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("UPDATE Report SET user_id = ?, title = ?, character_counts = ?, style = ?, language = ? WHERE id = ?")
	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}
	defer state.Close()
	id := context.Query("id")
	user_id := context.Query("user_id")
	title := context.Query("title")
	character_counts := context.Query("character_counts")
	style := context.Query("style")
	language := context.Query("language")
	_, err = state.Exec(user_id, title, character_counts, style, language, id)

	if err != nil {
		return "", fmt.Errorf("クエリ実行エラー: %v", err)
	}

	return "Report Updated", nil

}

func DeleteReport(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("DELETE FROM Report WHERE id = ?")

	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}

	defer state.Close()

	Id := context.Query("id")
	_, err = state.Exec(Id)

	if err != nil {
		return "", fmt.Errorf("クエリ実行エラー: %v", err)
	}

	return "Report Deleted", nil
}
