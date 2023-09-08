package controllers

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AlertController struct {
	service *service.AlertService
}

func NewAlertController() *AlertController {
	return &AlertController{
		service: service.NewAlertService(),
	}
}

func (ac *AlertController) GetAlertByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var alert models.Alert
	alert, err = ac.service.GetAlertByUserID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": alert,
	})
}

func (ac *AlertController) UpdateAlert(c *gin.Context) {
	req := new(dto.UpdateAlertRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"message": "invalid request",
			})
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	rawID, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id value not found",
		})
	}
	ID, _ := rawID.(uint)

	fmt.Println(req.EnableGif)
	if req.EnableGif != 1 {
		req.EnableGif = 0 // maybe temporary fix for can't send 0 value
	}
	err = ac.service.UpdateAlert(ID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "failed to update alert",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update alert",
	})
}

func (ac *AlertController) UpdateAlertDesign(c *gin.Context) {
	req := new(dto.UpdateAlertDesignRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	err = ac.service.UpdateAlertDesign(c.GetUint("ID"), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "failed to update alert design",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update alert design",
	})
}
