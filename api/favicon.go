package api

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// favicon.png is a copy of web/favicon.png — keep them in sync (using "make assets").
//
//go:embed favicon.png
var faviconPNG []byte

func handleFavicon(c *gin.Context) {
	c.Data(http.StatusOK, "image/png", faviconPNG)
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/favicon.ico", handleFavicon)
	})
}
