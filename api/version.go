package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/buildinfo"
)

type VersionEntity struct {
	Version      string `json:"version"`
	Commit       string `json:"commit"`
	BuildTime    string `json:"buildTime"`
	BuiltBy      string `json:"builtBy"`
	LastModified string `json:"lastModified,omitempty"`
}

// handleVersion godoc
//
//	@Summary	Version and resume last-modified date
//	@Tags		version
//	@Produce	json
//	@Success	200	{object}	VersionEntity
//	@Router		/version [get]
func handleVersion(c *gin.Context) {
	v := VersionEntity{
		Version:   buildinfo.Version,
		Commit:    buildinfo.Commit,
		BuildTime: buildinfo.BuildTime,
		BuiltBy:   buildinfo.BuiltBy,
	}
	if r, err := loadResume(); err == nil {
		v.LastModified = r.Meta.LastModified
	}
	c.JSON(http.StatusOK, v)
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/version", handleVersion)
	})
}
