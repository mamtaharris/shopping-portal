package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/shopping-portal/models"
)

func HomeHandler(ctx *gin.Context) {
	defaultResponse := models.HomeResponse{
		DefaultResponse: "Hello",
	}
	response, err := json.Marshal(defaultResponse)
	if err == nil {
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.String(http.StatusOK, string(response))
}
