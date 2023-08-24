package handler

import (
	"fmt"
	"net/http"

	"github.com/fowkeio/secjobs/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/ap1/v1

// @Summary		Show opening
// @Description	Show a new secjob opening
// @Tags			Openings
// @Accept			json
// @Produce		json
// @Param			id	query		string	true	"Opening identification"
// @Success		200	{object}	ShowOpeningResponse
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router			/opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, ErrorParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error showing opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
