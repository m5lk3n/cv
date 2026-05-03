package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type WorkEntity struct {
	Items []resume.Work `json:"items"`
}

type CompanyAchievements struct {
	Name       string   `json:"name"`
	Highlights []string `json:"highlights"`
}

type AchievementsEntity struct {
	Items []CompanyAchievements `json:"items"`
}

// handleWork godoc
//
//	@Summary	Work experience
//	@Tags		work
//	@Produce	json
//	@Success	200	{object}	WorkEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/work [get]
func handleWork(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, WorkEntity{Items: r.Work})
}

// handleAchievements godoc
//
//	@Summary	All work achievements grouped by company
//	@Tags		work
//	@Produce	json
//	@Success	200	{object}	AchievementsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/work/achievements [get]
//	@Router		/achievements [get]
func handleAchievements(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	items := make([]CompanyAchievements, 0, len(r.Work))
	for _, job := range r.Work {
		items = append(items, CompanyAchievements{
			Name:       job.Name,
			Highlights: job.Highlights,
		})
	}
	c.JSON(http.StatusOK, AchievementsEntity{Items: items})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/work", handleWork)
		e.GET("/work/achievements", handleAchievements)
		e.GET("/achievements", handleAchievements)
	})
}
