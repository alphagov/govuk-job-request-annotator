package init_app

import (
	"github.com/alphagov/govuk-job-request-annotator/routes"
	"github.com/gin-gonic/gin"
)

func InitGin(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)

	r := gin.New()

	routes.InitLogger(r)

	routes.InitRouter(r)

	return r
}
