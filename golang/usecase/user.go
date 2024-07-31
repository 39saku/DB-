package usecase

import (
	"database/sql"
	"db_assignment/domain"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(db *sql.DB, context *gin.Context) (string, int, error) {
	state, err := db.Prepare("INSERT INTO User (name) VALUES(?)")
	if err != nil {
		return "", 400, err
	}
	defer state.Close()

	Name := context.Query("name")

	if Name == "" {
		return "Query is wrong", 400, err
	}
	_, err = state.Exec(Name)
	if err != nil {
		return "", 400, err
	}

	return "User created", 200, nil
}
func GetUser(db *sql.DB, context *gin.Context) ([]domain.User, int, error) {
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		return nil, 400, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, 400, err
		}
		users = append(users, user)
	}

	return users, 200, nil
}
func GetUserId(db *sql.DB, context *gin.Context) ([]string, int, error) {

	Name := context.Query("name")
	rows, err := db.Query("SELECT id FROM User WHERE name = ?", Name)
	if err != nil {
		return nil, 400, err
	}
	var IDs []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, 400, err
		}
		IDs = append(IDs, id)
	}

	return IDs, 200, nil
}
func UpdateUser(db *sql.DB, context *gin.Context) (string, int, error) {
	state, err := db.Prepare("UPDATE User SET name = ? WHERE id = ?")
	if err != nil {
		return "", 400, err
	}
	defer state.Close()

	Id := context.Query("id")
	Name := context.Query("name")

	_, err = state.Exec(Name, Id)

	if err != nil {
		return "", 400, err
	}

	return "User updated", 200, nil
}
