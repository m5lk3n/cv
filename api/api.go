package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lttl.dev/cv/resume"
)

// loadResume is the indirection used by all handlers to load resume data.
// Tests override it to inject a fixture.
var loadResume = resume.Load

// ErrorResponse is the JSON shape returned for non-2xx responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

var routes []func(*gin.Engine)

func register(r func(*gin.Engine)) {
	routes = append(routes, r)
}

// NewRouter returns a Gin engine with all CV API routes registered.
func NewRouter() *gin.Engine {
	e := gin.New()
	e.SetTrustedProxies(nil) // disable warnings about trusting all proxies
	e.Use(gin.Recovery(), slogMiddleware())
	for _, r := range routes {
		r(e)
	}

	e.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.NoRoute(handleNotFound)
	return e
}

func handleNotFound(c *gin.Context) {
	slog.Warn("route not found",
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		"client_ip", c.ClientIP(),
	)
	c.JSON(http.StatusNotFound, ErrorResponse{Error: "not found"})
}

// Run starts the API server on the given address (e.g. ":8080").
func Run(addr string) error {
	return NewRouter().Run(addr)
}

// loadResumeOrError loads the resume; on failure it logs the error, writes a
// 500 JSON response, and returns ok=false so the caller can return early.
func loadResumeOrError(c *gin.Context) (resume.Resume, bool) {
	r, err := loadResume()
	if err != nil {
		slog.Error("load resume", "err", err, "path", c.FullPath())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return r, false
	}
	return r, true
}

func slogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		slog.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		)
	}
}
