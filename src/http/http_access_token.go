package http

import (
	"github.com/JenniO/bookstore_oauth-api/src/domain/access_token"
	accesstoken2 "github.com/JenniO/bookstore_oauth-api/src/services/access_token"
	"github.com/JenniO/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}
type accessTokenHandler struct {
	service accesstoken2.Service
}

func NewHandler(service accesstoken2.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusNotImplemented, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	at, err := h.service.Create(atr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
