package controllers

import (
	"cortify/common"
	"cortify/dao"
	"cortify/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//Kubeconfig

// func K8sConfig() {
// 	// Load the in-cluster config
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// Create the clientset
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// Service struct
type Service struct {
	serviceDAO dao.Service
}

// GetKnativeService by ID
func (s *Service) GetService(ctx *gin.Context) {
	var service models.Service
	var err error
	//name := ctx.Request.URL.Query()["name"][0]
	id := ctx.Param("_id")
	service, err = s.serviceDAO.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	} else {
		ctx.JSON(http.StatusOK, service)
		ctx.JSON(http.StatusOK, "Knative Service : ")
		// ctx.JSON(http.StatusOK, svc)
	}
}

// AddService function (Deploy)
func (s *Service) AddService(ctx *gin.Context) {

	// Define Data Model
	var service models.Service
	if err := ctx.BindJSON(&service); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Init Default Values
	service.ID = service.KService.Metadata.Name + "-" + service.KService.Metadata.Namespace + "-service"
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	//Construction
	// Construct Knative Service
	ksvc, err := dao.ConstructService(service.KService.Metadata.Name, service.KService.Metadata.Namespace, service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"Knative Service Cannot be constructed."})
		return
	}
	cm, err := dao.ConstructConfigMap(service.KService.Metadata.Name, service.KService.Metadata.Namespace, service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Message{"ConfigMap Cannot be constructed."})
		return
	}
	ctx.JSON(http.StatusOK, models.Message{"construction done."})

	// Create Knative Service
	err = dao.CreateService(service.KService.Metadata.Namespace, ksvc)

	//Create CM and Secrets
	if err == nil {
		err = dao.CreateConfigMap(service.KService.Metadata.Namespace, cm)
		if err == nil {
			err := s.serviceDAO.Insert(service)
			if err != nil {
				ctx.JSON(http.StatusForbidden, models.Error{common.StatusCodeUnknown, err.Error()})
				log.Debug("[ERROR]: DB - ", err)
			}
			ctx.JSON(http.StatusOK, models.Message{"ConfigMap created Successfully"})
		} else {
			ctx.JSON(http.StatusBadRequest, models.Message{"Failed to create ConfigMap."})
			log.Debug("[ERROR]: ConfigMap -  ", err)
		}
		ctx.JSON(http.StatusOK, models.Message{"Knative  Service created Successfully"})
	} else {
		ctx.JSON(http.StatusForbidden, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: Knative Service - ", err)

	}
}
