package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	RegisterRoutes(r)

	if err := r.Run(":81"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("hello world")

}
