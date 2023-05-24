package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAdHandler struct {
	findAllAds ad.FindAllAds
	findAdById ad.FindAdById
}

func NewGetAdHandler(dependencies *configuration.AdConfiguration) *GetAdHandler {
	return &GetAdHandler{
		findAdById: dependencies.FindAdByIdService,
		findAllAds: dependencies.FindAllAdsService,
	}
}

type JSONFindAllAdsResponse struct {
	Id string `json:"id"`
}

func (h GetAdHandler) FindAllAds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ads, err := h.findAllAds.Execute(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		switch {
		case len(ads) == 0:
			ctx.Status(http.StatusNoContent)
			return
		default:
			ctx.JSON(http.StatusOK, h.mapFindAllAdsToJSONResponse(ads))
			return
		}
	}
}

func (GetAdHandler) mapFindAllAdsToJSONResponse(allAdsResponse []ad.GetAdsDto) []JSONFindAllAdsResponse {
	jsonResponse := make([]JSONFindAllAdsResponse, 0)
	for _, response := range allAdsResponse {
		jsonResponse = append(jsonResponse, JSONFindAllAdsResponse{Id: response.Id})
	}
	return jsonResponse
}

type JSONFindAdByIdResponse struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Date        string  `json:"postedAt"`
}

func (h GetAdHandler) FindAdById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adResponse, err := h.findAdById.Execute(ctx, ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		switch {
		case len(adResponse.Id) == 0:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.JSON(http.StatusOK, h.mapFindAdByIdToJSONResponse(adResponse))
			return
		}
	}
}

func (GetAdHandler) mapFindAdByIdToJSONResponse(adByIdResponse ad.GetAdByIdDto) JSONFindAdByIdResponse {
	return JSONFindAdByIdResponse{
		Id:          adByIdResponse.Id,
		Title:       adByIdResponse.Title,
		Description: adByIdResponse.Description,
		Price:       adByIdResponse.Price,
		Date:        adByIdResponse.Date,
	}
}
