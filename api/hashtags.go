package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HashtagsEntity struct {
	Hashtags []string `json:"hashtags"`
}

// handleHashtags godoc
//
//	@Summary	Hashtags
//	@Tags		hashtags
//	@Produce	json
//	@Success	200	{object}	HashtagsEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/hashtags [get]
func handleHashtags(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, HashtagsEntity{Hashtags: r.XCV.Hashtags})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/hashtags", handleHashtags)
	})
}
