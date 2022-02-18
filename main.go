package main

import (
	// "fmt"
	// "math"

	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// func delay () {
// 	num := float64(1)
// 	for i := 0; i <= 100000000; i ++ {
// 		num += math.Sqrt(num)
// 	}
// 	fmt.Println(num)
// }

func simulateStress() {
  cmd := exec.Command("/usr/bin/stress-ng", "--iomix 8" )
	if err := cmd.Start(); err != nil {
		fmt.Printf("---- %v\n", err.Error())
	}
}

func main() {
	r := gin.Default()
	go simulateStress()
	r.GET("/ping", func(c *gin.Context) {
	  // delay()
		go simulateStress()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}