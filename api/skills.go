package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type SkillsEntity struct {
	Items []resume.Skill `json:"items"`
}

// handleSkills godoc
//
//	@Summary	Skills
//	@Tags		skills
//	@Produce	json
//	@Success	200	{object}	SkillsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/skills [get]
func handleSkills(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, SkillsEntity{Items: r.Skills})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/skills", handleSkills)
	})
}
