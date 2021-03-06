package controllers

import (
	"cortify/dao"
	"cortify/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Environment struct{}

func (e *Environment) CreateEnvironment(ctx *gin.Context) {

	var environment models.Environment
	if err := ctx.BindJSON(&environment); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	environment.ID = environment.Name + "-client"
	environment.CreatedAt = time.Now()

	// Construct Environment
	env := dao.ConstructEnvironment(environment)
	err := dao.CreateEnvironment(env)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, models.Message{"Environment Cannot be Created."})
		log.Debug("[ERROR]: Environment -", err)
		return
	}

	ctx.JSON(http.StatusOK, models.Message{"Environment created Successfully"})
	ctx.JSON(http.StatusOK, env)
}
