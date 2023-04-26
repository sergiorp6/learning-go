package postad

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application/postad"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"net/http"
)

func Handler(postAdService ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req httpPostAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id := ctx.Param("id")

		err := postAdService.Execute(NewRequest(id, req.Title, req.Description, req.Price))
		if err != nil {
			_ = fmt.Errorf("error posting ad: %s", id)
			if errors.Is(err, ErrTitleTooLong) {
				ctx.JSON(http.StatusBadRequest, err.Error())
			}
		}

		ctx.Status(http.StatusCreated)
	}
}

type httpPostAdRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
