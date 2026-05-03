package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type ProjectsEntity struct {
	Items []resume.Project `json:"items"`
}

// handleProjects godoc
//
//	@Summary	Projects
//	@Tags		projects
//	@Produce	json
//	@Success	200	{object}	ProjectsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/projects [get]
func handleProjects(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, ProjectsEntity{Items: r.Projects})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/projects", handleProjects)
	})
}
