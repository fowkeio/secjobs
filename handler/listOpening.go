package handler

import (
	"net/http"

	"github.com/fowkeio/secjobs/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/ap1/v1

// @Summary		List openings
// @Description	List all new secjob openings
// @Tags			Openings
// @Accept			json
// @Produce		json
// @Success		200	{object}	ListOpeningsResponse
// @Failure		500	{object}	ErrorResponse
// @Router			/openings [get]
func ListOpeninsgHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)

}
