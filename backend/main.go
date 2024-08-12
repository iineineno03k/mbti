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
	user1 := User{
		Id:        1,
		LastName:  "田中",
		FirstName: "太郎",
		Nickname:  "たなっち",
		MBTI:      constant.INTJ_A,
	}
	user2 := User{
		Id:        2,
		LastName:  "山田",
		FirstName: "花子",
		Nickname:  "はなちゃん",
		MBTI:      constant.ENFJ_T,
	}

	var users = []User{
		user1,
		user2,
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
	db.Close()

	// ルート設定
	e.GET("/user", getAllUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
