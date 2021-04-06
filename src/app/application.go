package app

import (
	"github.com/JenniO/bookstore_oauth-api/src/http"
	"github.com/JenniO/bookstore_oauth-api/src/repository/db"
	"github.com/JenniO/bookstore_oauth-api/src/repository/rest"
	accessToken2 "github.com/JenniO/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(accessToken2.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8084")
}
