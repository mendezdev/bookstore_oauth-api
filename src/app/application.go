package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mendezdev/bookstore_oauth-api/src/clients/cassandra"
	"github.com/mendezdev/bookstore_oauth-api/src/doamin/access_token"
	"github.com/mendezdev/bookstore_oauth-api/src/http"
	"github.com/mendezdev/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
