package main

import (
	"backend/constant"
	"fmt"
)

type User struct {
	LastName  string
	FirstName string
	Nickname  string
	MBTI      constant.MBTI
}

func main() {
	user := User{
		LastName:  "Suzuki",
		FirstName: "Taro",
		Nickname:  "Suzu",
		MBTI:      constant.INTJ_A,
	}

	fmt.Printf(user.Nickname)
}
