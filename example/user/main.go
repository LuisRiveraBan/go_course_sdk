package main

import (
	"errors"
	"fmt"
	userSdk "github.com/LuisRiveraBan/go_course_sdk/user"
	"os"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("f5393bae-5a78-478f-811d-766d03184a8b")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)

}
