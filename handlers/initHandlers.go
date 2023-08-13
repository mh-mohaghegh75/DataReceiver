package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {

	db := GetConnection(dsn)

	err := db.AutoMigrate(&User{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("User Migrate Error: %s", err))
	}

	err = db.AutoMigrate(&Data{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("User Migrate Error: %s", err))
	}
}
