package findbyid

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/findbyid"
	"net/http"
)

func FindByIdHandler(findByIdService FindByIdService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		ad, err := findByIdService.Execute(NewFindByIdRequest(id))

		if err != nil {
			_ = fmt.Errorf("error fetching ad: %s", id)
		}

		if ad != nil {
			response := httpAdResponse{
				ID:          ad.Id().String(),
				Title:       ad.Title().Value(),
				Description: ad.Description().Value(),
				Price:       ad.Price().Value(),
			}
			ctx.JSON(http.StatusOK, response)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	}
}

type httpAdResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
