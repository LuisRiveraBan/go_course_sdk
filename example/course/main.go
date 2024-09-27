package main

import (
	"errors"
	"fmt"
	courseSdk "github.com/LuisRiveraBan/go_course_sdk/course"
	"os"
)

func main() {

	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("ff5810c2-55b3-4d0a-8dd2-8de9b8486d8c")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(course)
}
