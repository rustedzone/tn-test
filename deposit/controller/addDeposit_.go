package depositcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	service "tn-test/deposit/service"

	"gopkg.in/gin-gonic/gin.v1"
)

func AddDeposit_(c *gin.Context) {

	//params handler
	var request map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(c.Request.Body)
		log.Println("error on parsing request")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	log.Println("request :", request)

	err = service.AddDeposit(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}
