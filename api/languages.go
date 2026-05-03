package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type LanguagesEntity struct {
	Items []resume.Language `json:"items"`
}

// handleLanguages godoc
//
//	@Summary	Languages
//	@Tags		languages
//	@Produce	json
//	@Success	200	{object}	LanguagesEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/languages [get]
func handleLanguages(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, LanguagesEntity{Items: r.Languages})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/languages", handleLanguages)
	})
}
