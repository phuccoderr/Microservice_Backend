package controller

import (
	"github.com/gin-gonic/gin"
	"mail-service/internal/dto"
	"mail-service/internal/service"
	"net/http"
)

type MailController struct {
	mailService service.IMailService
}

func NewMailController(mailService service.IMailService) *MailController {
	return &MailController{
		mailService: mailService,
	}
}

func (mc *MailController) SendFeedBack(c *gin.Context) {
	var feedback dto.Feedback
	err := c.ShouldBindJSON(&feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error json": err.Error()})
		return
	}
	err = mc.mailService.Feedback(feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error json": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
