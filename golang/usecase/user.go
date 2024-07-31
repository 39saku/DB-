package usecase

import (
	"database/sql"
	"db_assignment/domain"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(db *sql.DB, context *gin.Context) {
	state, err := db.Prepare("INSERT INTO User (name) VALUES(?)")
	if err != nil {
		context.JSON(400, err)
	}
	defer state.Close()

	Name := context.Query("name")

	_, err = state.Exec(Name)
	if err != nil {
		context.JSON(400, err)
	}
	context.JSON(200, gin.H{"message": "User created"})
}
func GetUser(db *sql.DB, context *gin.Context) {
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		context.JSON(400, err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			context.JSON(400, err)
		}
		users = append(users, user)
	}

	context.JSON(200, users)
}
func GetUserId(db *sql.DB, context *gin.Context) {

	Name := context.Query("name")
	rows, err := db.Query("SELECT id FROM User WHERE name = ?", Name)
	if err != nil {
		context.JSON(400, err)
	}
	var IDs []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			context.JSON(400, err)
		}
		IDs = append(IDs, id)
	}

	context.JSON(200, IDs)
}
func UpdateUser(db *sql.DB, context *gin.Context) {
	state, err := db.Prepare("UPDATE User SET name = ? WHERE id = ?")
	if err != nil {
		context.JSON(400, err)
	}
	defer state.Close()

	Id := context.Query("id")
	Name := context.Query("name")

	_, err = state.Exec(Name, Id)

	if err != nil {
		context.JSON(400, err)
	}

	context.JSON(200, gin.H{"message": "User updated"})
}
