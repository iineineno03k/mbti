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
	LastName  string        `json:"lastName" form:"lastName" validate:"required"`
	FirstName string        `json:"firstName" form:"firstName" validate:"required"`
	Nickname  string        `json:"nickname" form:"nickname"`
	MBTI      constant.MBTI `json:"mbti" form:"mbti" validate:"required"`
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

func postUser(c echo.Context) error {
	db, ok := c.Get("db").(*sql.DB)
	if !ok || db == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database connection is not available")
	}

	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// データベースにユーザーを挿入
	query := `INSERT INTO mbti_user (last_name, first_name, nickname, mbti) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, u.LastName, u.FirstName, u.Nickname, u.MBTI)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, u)
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
	e.POST("/user", postUser)

	e.Logger.Fatal(e.Start(":8080"))

}
