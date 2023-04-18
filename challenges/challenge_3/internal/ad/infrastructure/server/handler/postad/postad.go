package postad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/postad"
	"net/http"
	"time"
)

func PostAdHandler(postAdService PostAdService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req httpPostAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id := ctx.Param("id")

		err := postAdService.Execute(NewPostAdRequest(id, req.Title, req.Description, req.Price, time.Now()))
		if err != nil {
			_ = fmt.Errorf("error posting ad: %s", id)
		}

		ctx.Status(http.StatusCreated)
	}
}

type httpPostAdRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
