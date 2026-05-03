package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type ReferencesEntity struct {
	Items []resume.Reference `json:"items"`
}

// handleReferences godoc
//
//	@Summary	References / recommendations
//	@Tags		references
//	@Produce	json
//	@Success	200	{object}	ReferencesEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/references [get]
func handleReferences(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, ReferencesEntity{Items: r.References})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/references", handleReferences)
	})
}
