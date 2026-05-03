package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type EducationEntity struct {
	Items []resume.Education `json:"items"`
}

// handleEducation godoc
//
//	@Summary	Education history
//	@Tags		education
//	@Produce	json
//	@Success	200	{object}	EducationEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/education [get]
func handleEducation(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, EducationEntity{Items: r.Education})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/education", handleEducation)
	})
}
