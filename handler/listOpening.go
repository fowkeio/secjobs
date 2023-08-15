package handler

import (
	"net/http"

	"github.com/fowkeio/secjobs/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeninsgHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)

}
