package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lttl.dev/cv/resume"
)

type CertificatesEntity struct {
	Items []resume.Certificate `json:"items"`
}

// handleCertificates godoc
//
//	@Summary	Certificates
//	@Tags		certificates
//	@Produce	json
//	@Success	200	{object}	CertificatesEntity
//	@Failure	500	{object}	ErrorResponse
//	@Router		/certificates [get]
func handleCertificates(c *gin.Context) {
	r, ok := loadResumeOrError(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, CertificatesEntity{Items: r.Certificates})
}

func init() {
	register(func(e *gin.Engine) {
		e.GET("/certificates", handleCertificates)
	})
}
