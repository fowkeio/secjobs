package handler

import (
	"fmt"
	"net/http"

	"github.com/fowkeio/secjobs/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/ap1/v1

// @Summary		Delete opening
// @Description	Delete a new secjob opening
// @Tags			Openings
// @Accept			json
// @Produce		json
// @Param			id	query		string	true	"Opening identification"
// @Success		200	{object}	DeleteOpeningResponse
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router			/opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, ErrorParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
