package main

// 依存関係のインポート
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	Id   string
	Name string
}
type Report struct {
	Id               string `json:"id"`
	User_id          string `json:"user_id"`
	Title            string `json:"title"`
	Character_counts int    `json:"character_counts"`
	Style            int    `json:"style"`
	Language         int    `json:"language"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go [import|export]")
		os.Exit(1)
	}

	// sys.argv[1]でコマンドライン引数を取得
	// command := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ここでDB接続情報を環境変数から取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	var db *sql.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// 接続再試行用のカウント．30カウントしたら失敗とする．設定する理由は接続情報はあっているのに1回でつながらない場合がよくあるからである
	for i := 0; i < 30; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}
	defer db.Close()

	// テーブルを作成するSQLクエリ(テーブルが存在しない場合のみ実行)ついでに実行(exec)
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS User (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
     )
     `)
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Report (
        id INT AUTO_INCREMENT PRIMARY KEY,
		User_id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
		character_counts INT NOT NULL,
		style INT NOT NULL,
		language INT NOT NULL
	)
    `)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("テーブルが作成されました")

	r := gin.Default()
	r.GET("/create_user", func(c *gin.Context) {
		str := create_user(db, c)
		println(str)
	})

	r.GET("/get_user", func(c *gin.Context) {
		str, _ := get_user(db)
		c.JSON(200, str)
	})
	r.GET("/get_user_id", func(c *gin.Context) {
		str, _ := get_user_id(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/update_user", func(c *gin.Context) {
		str := update_user(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/get_a_list_report", func(c *gin.Context) {
		reports, err := get_a_list_report(db, c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, reports)
	})
	r.GET("/get_report", func(c *gin.Context) {
		str, _ := get_report(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/create_report", func(c *gin.Context) {
		str, _ := create_report(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/update_report", func(c *gin.Context) {
		str, _ := update_report(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/delete_report", func(c *gin.Context) {
		str, _ := delete_report(db, c)
		c.JSON(200, gin.H{
			"message": str,
		})
	})

	r.Run(":8080")
}

func create_user(db *sql.DB, context *gin.Context) string {
	state, err := db.Prepare("INSERT INTO User (name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer state.Close()

	user := User{
		Id:   "1",
		Name: context.Query("name"),
	}
	_, err = state.Exec(user.Name)
	if err != nil {
		log.Fatal(err)
	}
	return "User created"
}
func get_user(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func get_user_id(db *sql.DB, context *gin.Context) ([]string, error) {

	user := User{
		Id:   "1",
		Name: context.Query("name"),
	}
	rows, err := db.Query("SELECT id FROM User WHERE name = ?", user.Name)
	if err != nil {
		log.Fatal(err)
	}
	var IDs []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		IDs = append(IDs, id)
	}
	return IDs, nil
}
func update_user(db *sql.DB, context *gin.Context) string {
	state, err := db.Prepare("UPDATE User SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer state.Close()
	user := User{
		Id:   context.Query("id"),
		Name: context.Query("name"),
	}
	_, err = state.Exec(user.Name, user.Id)

	return "User Updated"
}

func get_a_list_report(db *sql.DB, context *gin.Context) ([]Report, error) {
	userID := context.Query("id")
	rows, err := db.Query("SELECT * FROM Report WHERE User_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー: %v", err)
	}
	defer rows.Close()
	var reports []Report
	for rows.Next() {
		var report Report
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

func get_report(db *sql.DB, context *gin.Context) ([]Report, error) {
	report := Report{
		Id:               context.Query("id"),
		User_id:          "a",
		Title:            "a",
		Character_counts: 1,
		Style:            1,
		Language:         1,
	}
	rows, err := db.Query("SELECT * FROM Report WHERE id = ?", report.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reports []Report
	for rows.Next() {
		var report Report
		err := rows.Scan(&report.Id, &report.User_id, &report.Title, &report.Character_counts, &report.Style, &report.Language)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func create_report(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("INSERT INTO Report (User_id, title, character_counts, style, language) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("can not prepare: %v", err)
	}
	defer state.Close()

	User_id := context.Query("User_id")
	title := context.Query("title")
	characterCounts := context.Query("character_counts")
	style := context.Query("style")
	if err != nil {
		return "", fmt.Errorf("invalid style: %v", err)
	}
	language := context.Query("language")
	if err != nil {
		return "", fmt.Errorf("invalid language: %v", err)
	}

	_, err = state.Exec(User_id, title, characterCounts, style, language)
	if err != nil {
		return "", fmt.Errorf("can not Exec: %v", err)
	}
	return "report created", nil
}

func update_report(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("UPDATE Report SET User_id = ?, title = ?, character_counts = ?, style = ?, language = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer state.Close()
	Id := context.Query("Id")
	User_id := context.Query("User_id")
	Title := context.Query("Title")
	Character_counts := context.Query("Character_counts")
	Style := context.Query("Style")
	Language := context.Query("Language")
	_, err = state.Exec(User_id, Title, Character_counts, Style, Language, Id)

	return "Report Updated", nil

}

func delete_report(db *sql.DB, context *gin.Context) (string, error) {
	state, err := db.Prepare("DELETE FROM Report WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer state.Close()

	Id := context.Query("Id")
	_, err = state.Exec(Id)

	return "Report Deleted", nil
}
