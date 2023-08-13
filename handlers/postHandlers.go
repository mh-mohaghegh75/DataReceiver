package handler

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func InsertData(c *gin.Context) {
	db := GetConnection(dsn)

	//Recieve Data from request body
	var v Data
	err := c.ShouldBindJSON(&v)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Bind Data Error: %s", err))
		return
	}

	//Check for duplication
	var temp Data
	row := db.Raw("SELECT * FROM data WHERE id=?", v.ID).Scan(&temp).RowsAffected
	if row != 0 {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("duplication Error: %s", err))
		return
	}

	// Get users used quotas
	var usedQ UsedQouta
	err = db.Raw("SELECT COUNT(*) AS count , user_id FROM data WHERE user_id=? GROUP BY user_id", v.UserID).Scan(&usedQ).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Used Quota Error: %s", err))
		return
	}

	// Get users quota
	userQuota, err := GetQuota(v.UserID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Get Quota Error: %s", err))
		return
	}

	// Checking whether used quotas reached maximum
	err = QuotaCheck(usedQ, userQuota)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Quota Check Error: %s", err))
		return
	} else {
		err = db.Exec("INSERT INTO data(id, user_id) VALUES(?, ?)", v.ID, v.UserID).Error
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Insertion Error: %s", err))
			return
		}

		c.JSON(http.StatusAccepted, "Data send to queue")
		// SendQueue() //Send data to other services
	}
}

// Get user quota
func GetQuota(ui UserId) (int, error) {
	db := GetConnection(dsn)

	var result User

	err := db.Raw(`SELECT * FROM "user" WHERE id = ?`, ui).Scan(&result).Error
	if err != nil {
		return -1, err
	}
	return int(result.Quota), nil
}

// Check for reaching max Quota
func QuotaCheck(q UsedQouta, quota int) error {
	if q.Count >= quota {
		return fmt.Errorf("User qouta's reached maximum")
	}
	return nil
}

//Add data to temp queue, slice mode temp queue
func AddTempQueue(queue []byte, data byte) []byte {
	queue = append(queue, data)
	fmt.Println("Data added to temp queue.")
	return queue
}

//remove data from temp queue
func RemoveTempQueue(queue []byte, index uint) []byte {
	var tqueue []byte
	for i, v := range queue {
		if i == int(index) {
			continue
		}
		tqueue = append(tqueue, v)
	}
	return tqueue
}

//Data will be sent to other services
func SendQueue() {
	for {
		//send data to queue and remove from temp queue using two function above
	}
}
