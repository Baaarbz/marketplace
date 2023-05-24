package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/pkg/application/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveAdHandler struct {
	saveAd ad.SaveAd
}

func NewSaveAdHandler(dependencies *configuration.AdConfiguration) *SaveAdHandler {
	return &SaveAdHandler{
		saveAd: dependencies.SaveAdService,
	}
}

type JSONSaveAdResponse struct {
	Id string `json:"id"`
}

type JSONSaveAdRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
}

func (h SaveAdHandler) SaveAd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req JSONSaveAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if savedAd, err := h.saveAd.Execute(ctx, h.mapJSONSaveAdToRequest(req)); err != nil {
			switch {
			case errors.Is(err, ErrAdIdBadFormat) ||
				errors.Is(err, ErrTitleBadFormat) ||
				errors.Is(err, ErrDescriptionBadFormat) ||
				errors.Is(err, ErrPriceBadFormat):
				ctx.String(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			ctx.JSON(http.StatusCreated, h.mapSaveAdToJsonResponse(savedAd))
		}
	}
}

func (SaveAdHandler) mapSaveAdToJsonResponse(response ad.SaveAdDtoResponse) JSONSaveAdResponse {
	return JSONSaveAdResponse{
		Id: response.Id,
	}
}

func (SaveAdHandler) mapJSONSaveAdToRequest(request JSONSaveAdRequest) ad.SaveAdDtoRequest {
	return ad.SaveAdDtoRequest{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
	}
}
