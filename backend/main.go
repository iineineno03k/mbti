package main

import (
	"backend/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id        int
	LastName  string
	FirstName string
	Nickname  string
	MBTI      constant.MBTI
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
	e.GET("/user", getAllUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
