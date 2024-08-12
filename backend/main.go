package main

import (
	"backend/constant"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int           `json:"id"`
	LastName  string        `json:"lastName"`
	FirstName string        `json:"firstName"`
	Nickname  string        `json:"nickname"`
	MBTI      constant.MBTI `json:"mbti"`
}

func getAllUsers(c echo.Context) error {
	db, ok := c.Get("db").(*sql.DB)
	if !ok || db == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database connection is not available")
	}

	rows, err := db.Query("SELECT id, last_name, first_name, nickname, mbti FROM mbti_user")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "取得に失敗しました。")
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.LastName, &user.FirstName, &user.Nickname, &user.MBTI)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	// db接続
	godotenv.Load(".env")
	connStr := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USERNAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// ミドルウェアを使用してリクエストごとにデータベース接続を注入
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// ルート設定
	e.GET("/user", getAllUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
