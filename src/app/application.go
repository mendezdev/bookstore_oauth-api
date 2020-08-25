package app

import (
	"github.com/gin-gonic/gin"

	"github.com/mendezdev/bookstore_oauth-api/src/http"
	"github.com/mendezdev/bookstore_oauth-api/src/repository/db"
	"github.com/mendezdev/bookstore_oauth-api/src/repository/rest"
	"github.com/mendezdev/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

// StartApplication is the main starting point for the application
func StartApplication() {
	//atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	atHandler := http.NewHandler(
		access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
