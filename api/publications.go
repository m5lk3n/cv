package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type PublicationsEntity struct {
	Items []resume.Publication `json:"items"`
}

// handlePublications godoc
//
//	@Summary	Publications
//	@Tags		publications
//	@Produce	json
//	@Success	200	{object}	PublicationsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/publications [get]
func handlePublications(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, PublicationsEntity{Items: r.Publications})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/publications", handlePublications)
	})
}
