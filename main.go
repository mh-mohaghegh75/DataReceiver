package main

import (
	handler "entrance/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//v1 stand for version one
	v1 := router.Group("/")
	{
		//This API made for table creation and test
		v1.GET("/migrate", handler.Migrate)

		create := v1.Group("/")
		{
			create.POST("/data", handler.InsertData)
		}
		delete := v1.Group("/")
		{
			delete.DELETE("/data/:id", handler.DeleteData)
		}
	}

	router.Run("localhost:1020")
}
