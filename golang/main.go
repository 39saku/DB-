package main

import (
	"database/sql"
	"db_assignment/handler"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go [import|export]")
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS User (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
     )
     `)
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Report (
        id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
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
		handler.CreateUserHandler(db, c)
	})

	r.GET("/get_user", func(c *gin.Context) {
		handler.GetUserHangler(db, c)
	})
	r.GET("/get_user_id", func(c *gin.Context) {
		handler.GetUserIdHandler(db, c)
	})
	r.GET("/update_user", func(c *gin.Context) {
		handler.UpdateUserHandler(db, c)
	})
	r.GET("/get_a_list_report", func(c *gin.Context) {
		handler.GetListReportHandler(db, c)
	})
	r.GET("/get_report", func(c *gin.Context) {
		handler.GetReportHandler(db, c)
	})
	r.GET("/create_report", func(c *gin.Context) {
		handler.CreateReportHandler(db, c)
	})
	r.GET("/update_report", func(c *gin.Context) {
		handler.UpdateReportHandler(db, c)
	})
	r.GET("/delete_report", func(c *gin.Context) {
		handler.DeleteReportHandler(db, c)
	})
	r.Run(":8080")
}
