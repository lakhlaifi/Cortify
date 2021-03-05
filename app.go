package main

import (
	"cortify/common"
	"cortify/controllers"
	"cortify/db"
	"io"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Main struct
type Main struct {
	router *gin.Engine
}

// init get Called before the main
func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize Database
	err = db.Database.Init()
	if err != nil {
		return err
	}
	// Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()
	return nil
}

func main() {
	m := Main{}
	// Initialize server
	if m.initServer() != nil {
		return
	}
	defer db.Database.Close()
	m.router.Use(cors.Default())

	// Define controller for Services
	s := controllers.Service{}
	e := controllers.Environment{}
	// Version API
	v1 := m.router.Group("/api/v1")
	{
		//Services
		// Add new KSVC
		v1.POST("/service", s.CreateService)
		v1.POST("/environment", e.CreateEnvironment)

		// Get One KSVC
		// v1.GET("/service/:_id", c.GetService)
	}
	m.router.NoRoute(func(c *gin.Context) {
		// In Gin Response are
		c.JSON(404, gin.H{"message": "Not Found"})
	})
	m.router.Run(common.Config.Port)

}
