package api

import (
	"agnos/backend/src/mapper"
	"agnos/backend/src/usecase/password"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	PasswordUsecase password.Usecase
}

func NewAPIHandler(passwordUsecase password.Usecase) *APIHandler {
	return &APIHandler{PasswordUsecase: passwordUsecase}
}

func (h *APIHandler) StrongPasswordSteps(c *gin.Context) {
	var reqBody map[string]interface{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	request := mapper.MapToPasswordRequest(reqBody)
	response, err := h.PasswordUsecase.CalculateNumOfSteps(request)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, response)
	}
}
