package main

import (
	"github.com/darwin1224/go-test/config"
	"github.com/darwin1224/go-test/controllers"
	"github.com/darwin1224/go-test/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.InitDB()
	articleController := &controllers.ArticleController{Service: &services.ArticleService{Repo: db}}
	r := gin.Default()
	r.GET("/", articleController.Index)
	r.GET("/:id", articleController.Show)
	r.POST("/", articleController.Store)
	r.PUT("/:id", articleController.Update)
	r.DELETE("/:id", articleController.Destroy)
	r.Run(":8000")
}
