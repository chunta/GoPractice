package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	r := gin.Default()

	Register(r)

	if err := r.Run(":80"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("hello world")
}
