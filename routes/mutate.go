package routes

import (
	"io"
	"net/http"

	m "github.com/alphagov/govuk-job-request-annotator/pkg/mutate"
	"github.com/alphagov/govuk-job-request-annotator/utils"
	"github.com/gin-gonic/gin"
)

func initMutatePod(r *gin.Engine) {
	r.POST("/mutate/configmap", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()

		if err != nil {
			errObj := utils.Response{
				Status: http.StatusInternalServerError,
				Data:   nil,
			}

			utils.SendResponse(c, errObj)
		}

		mutated, err := m.Mutate(body)
		if err != nil {
			errObj := utils.Response{
				Status: http.StatusInternalServerError,
				Data:   nil,
			}

			utils.SendResponse(c, errObj)
		}

		_, writeErr := c.Writer.Write(mutated)
		if writeErr != nil {
			errObj := utils.Response{
				Status: http.StatusInternalServerError,
				Data:   nil,
			}

			utils.SendResponse(c, errObj)
		}
	})
}
