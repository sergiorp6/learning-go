package getadslist

import (
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/getadslist"
	"net/http"
	"strconv"
)

func Handler(getAdsListService ServiceInterface) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		limitStr := ctx.DefaultQuery("limit", "10")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be a number"})
			return
		}
		ads := getAdsListService.Execute(NewRequest(limit))
		var adsResponse []httpAdResponse

		for _, ad := range ads {
			adsResponse = append(
				adsResponse,
				httpAdResponse{
					ID:          ad.Id().String(),
					Title:       ad.Title().Value(),
					Description: ad.Description().Value(),
					Price:       ad.Price().Value(),
				},
			)
		}
		ctx.JSON(http.StatusOK, httpAdListResponse{adsResponse})
	}
}

type httpAdListResponse struct {
	Data []httpAdResponse `json:"data"`
}
type httpAdResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
