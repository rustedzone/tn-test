package depositcontroller

import (
	"net/http"

	service "tn-test/deposit/service"

	"gopkg.in/gin-gonic/gin.v1"
)

func GetDepositAccount_(c *gin.Context) {

	//params handler
	account := c.Param("account")

	data, err := service.GetDepositAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": data})
}
