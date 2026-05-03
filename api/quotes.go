package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type QuotesEntity struct {
	Items []resume.Quote `json:"items"`
}

// handleQuotes godoc
//
//	@Summary	Quotes
//	@Tags		quotes
//	@Produce	json
//	@Success	200	{object}	QuotesEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/quotes [get]
func handleQuotes(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, QuotesEntity{Items: r.XCV.Quotes})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/quotes", handleQuotes)
	})
}
