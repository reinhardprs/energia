package controller

import (
	"net/http"

	"energia/controller/suggestion/request"
	"energia/controller/suggestion/response"
	"energia/service/suggestion"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type SuggestionController struct {
	suggestionService suggestion.SuggestionServiceInterface
}

func NewSuggestionController(ss suggestion.SuggestionServiceInterface) *SuggestionController {
	return &SuggestionController{
		suggestionService: ss,
	}
}

// GetSuggestions is the controller for the get suggestions endpoint
// @Summary Get suggestions
// @Description Get suggestions
// @Tags Suggestion
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param city body request.SuggestionRequest true "City to get suggestions"
// @Success 200 {object} response.SuggestionResponse
// @Failure 400 {object} base.BaseResponse
// @Router /suggestion [get]
func (c *SuggestionController) GetSuggestions(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	var req request.SuggestionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal membaca input"})
	}
	if req.City == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Nama kota harus diberikan"})
	}

	suggestionEntity, err := c.suggestionService.GetSuggestion(ctx.Request().Context(), userID, req.City)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	res := response.FromEntities(suggestionEntity)
	return ctx.JSON(http.StatusOK, res)
}
