package controllers

import (
	"net/http"
	"strconv"

	"github.com/darwin1224/go-test/models"

	"github.com/darwin1224/go-test/services"
	"github.com/gin-gonic/gin"
)

// ArticleController -- controller instance
type ArticleController struct {
	Service *services.ArticleService
}

// Index -- Get all resources in storage
func (a *ArticleController) Index(c *gin.Context) {
	articles := a.Service.GetAllArticle()
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

// Show -- Get a single resource in storage
func (a *ArticleController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	article := a.Service.GetArticleByID(&id)
	c.JSON(http.StatusOK, article)
}

// Store -- Insert a single resource from storage
func (a *ArticleController) Store(c *gin.Context) {
	var article models.Article
	c.Bind(&article)
	a.Service.InsertArticle(&article)
	c.JSON(http.StatusCreated, article)
}

// Update -- Update a single resource from storage
func (a *ArticleController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	article := a.Service.GetArticleByID(&id)
	c.Bind(&article)
	a.Service.UpdateArticle(&article)
	c.JSON(http.StatusOK, article)
}

// Destroy -- Delete a single resource from storage
func (a *ArticleController) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	article := a.Service.GetArticleByID(&id)
	a.Service.DeleteArticle(&article)
	c.JSON(http.StatusOK, article)
}
