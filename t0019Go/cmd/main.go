package main

import "github.com/sharin-sushi/0019testGin.git/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
