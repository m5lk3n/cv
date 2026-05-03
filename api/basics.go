package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type BasicsEntity struct {
	Name     string           `json:"name"`
	Label    string           `json:"label"`
	Image    string           `json:"image,omitempty"`
	Email    string           `json:"email,omitempty"`
	Phone    string           `json:"phone,omitempty"`
	URL      string           `json:"url,omitempty"`
	Summary  string           `json:"summary,omitempty"`
	Location resume.Location  `json:"location"`
	Profiles []resume.Profile `json:"profiles"`
}

// handleBasics godoc
//
//	@Summary	Basics including contact info
//	@Tags		basics
//	@Produce	json
//	@Success	200	{object}	BasicsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/basics [get]
func handleBasics(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	b := r.Basics
	c.JSON(http.StatusOK, BasicsEntity{
		Name:     b.Name,
		Label:    b.Label,
		Image:    b.Image,
		Email:    b.Email,
		Phone:    b.Phone,
		URL:      b.URL,
		Summary:  b.Summary,
		Location: b.Location,
		Profiles: b.Profiles,
	})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/basics", handleBasics)
	})
}
