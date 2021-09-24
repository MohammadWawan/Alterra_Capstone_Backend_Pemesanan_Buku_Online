package main

import (
	"alterra/configs"
	"alterra/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
