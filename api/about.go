package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AboutEntity struct {
	Summary string `json:"summary"`
}

// handleAbout godoc
//
//	@Summary	About summary
//	@Tags		about
//	@Produce	json
//	@Success	200	{object}	AboutEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/about [get]
func handleAbout(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, AboutEntity{Summary: r.Basics.Summary})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/about", handleAbout)
	})
}
