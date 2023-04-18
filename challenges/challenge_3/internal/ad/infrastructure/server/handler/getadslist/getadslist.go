package findbyid

import (
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/getadslist"
	"net/http"
)

func GetAdsListHandler(getAdsListService GetAdsListService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ads := getAdsListService.Execute(NewGetAdsListRequest(5))
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
