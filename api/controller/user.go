package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yurianxdev/rest-example/api/model"
	"github.com/yurianxdev/rest-example/api/repository"
)

var repo repository.UserRepository

func InitRepository(r repository.UserRepository) {
	repo = r
}

func GETUsers(c *gin.Context) {
	users, err := repo.ListUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, users)
}

func POSTUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err = repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user on database:\n%v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, user)
}
