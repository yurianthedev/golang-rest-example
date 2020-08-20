package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yurianxdev/rest-example/api/model"
	"github.com/yurianxdev/rest-example/api/repository/postgres"
)

var repository = postgres.UserRepository{}

func GETUsers(c *gin.Context) {
	users, err := repository.ListUsers()
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

	user, err = repository.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user on database:\n%v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, user)
}
