package controllers

import (
	"cortify/dao"
	"cortify/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	serviceDAO dao.Service
}

// CreateService function (Deploy)
func (s *Service) CreateService(ctx *gin.Context) {

	// Define Data Model
	var service models.Service
	if err := ctx.BindJSON(&service); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Init Default Values
	service.ID = service.Name + "-" + service.Namespace + "-service"
	service.CreatedAt = time.Now()

	// Construct Knative Service
	ksvc, err := dao.ConstructService(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"Knative Service Cannot be constructed."})
		log.Debug("[ERROR]: Construction KSVC -", err)
		return
	}
	// Construct ConfigMap
	cm, err := dao.ConstructConfigMap(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"ConfigMap Cannot be constructed."})
		log.Debug("[ERROR]: Construction CM -", err)
		return
	}
	// Construct Secret
	secret, err := dao.ConstructSecret(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"Secret Cannot be constructed."})
		log.Debug("[ERROR]: Construction Secret -", err)
		return
	}

	//TODO Dependencies Graph when creating Resources.
	// Creation with no Handled exception
	err = dao.CreateSecret(secret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"Secret Cannot be Created."})
	}
	err = dao.CreateConfigMap(cm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"ConfigMap Cannot be Created."})
	}
	err = dao.CreateService(ksvc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"Knative Service Cannot be Created."})
	}

	ctx.JSON(http.StatusOK, models.Message{"Service created Successfully"})
	ctx.JSON(http.StatusOK, service)
}
