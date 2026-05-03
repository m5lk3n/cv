package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type VolunteeringEntity struct {
	Items []resume.Volunteer `json:"items"`
}

// handleVolunteering godoc
//
//	@Summary	Volunteering experience
//	@Tags		volunteering
//	@Produce	json
//	@Success	200	{object}	VolunteeringEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/volunteering [get]
func handleVolunteering(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, VolunteeringEntity{Items: r.Volunteer})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/volunteering", handleVolunteering)
	})
}
