package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete Handler for deletion
func DeleteData(c *gin.Context) {
	db := GetConnection(dsn)

	id := c.Param("id")

	err := db.Exec("DELETE FROM data WHERE id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Deletion Error: %s", err))
	}
	c.Status(http.StatusNoContent)
}
