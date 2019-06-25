package main

import "github.com/iceyang/gin-exam/pkg/router"

func main() {
	engine := router.Default()
	engine.Run()
}
