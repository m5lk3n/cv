package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type FAQsEntity struct {
	Items []resume.FAQ `json:"items"`
}

// handleFAQs godoc
//
//	@Summary	FAQs
//	@Tags		faqs
//	@Produce	json
//	@Success	200	{object}	FAQsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/faqs [get]
func handleFAQs(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, FAQsEntity{Items: r.XCV.FAQs})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/faqs", handleFAQs)
	})
}
