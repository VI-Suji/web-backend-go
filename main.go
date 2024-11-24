package main

import (
	"fmt"

	"github.com/VI-Suji/web-backend-go/handlers"
	"github.com/VI-Suji/web-backend-go/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skill map")
	r := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManager)
	userHandler.RegisterUserApis(r)

	r.Run(":3001")
}
