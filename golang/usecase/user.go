package usecase

import (
	"database/sql"
	"db_assignment/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("INSERT INTO User (name) VALUES(?)")
	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}
	defer state.Close()

	Name := context.Query("name")

	_, err = state.Exec(Name)
	if err != nil {
		return "", fmt.Errorf("クエリ実行エラー: %v", err)
	}
	return "User created", err
}
func GetUser(db *sql.DB) ([]domain.User, error) {
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		return nil, fmt.Errorf("can not prepare: %v", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("行のスキャンエラー: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("行の反復中のエラー: %v", err)
	}

	return users, nil
}
func GetUserId(db *sql.DB, context *gin.Context) ([]string, error) {

	Name := context.Query("name")
	rows, err := db.Query("SELECT id FROM User WHERE name = ?", Name)
	if err != nil {
		return nil, fmt.Errorf("can not prepare: %v", err)
	}
	var IDs []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("行のスキャンエラー: %v", err)
		}
		IDs = append(IDs, id)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("行の反復中のエラー: %v", err)
	}

	return IDs, nil
}
func UpdateUser(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("UPDATE User SET name = ? WHERE id = ?")
	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}
	defer state.Close()

	Id := context.Query("id")
	Name := context.Query("name")

	_, err = state.Exec(Name, Id)

	if err != nil {
		return "", fmt.Errorf("クエリ実行エラー: %v", err)
	}

	return "User Updated", nil
}
