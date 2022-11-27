package api

import "github.com/gin-gonic/gin"

type Server struct{}

func (s *Server) Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/urlshort", s.urlshortener)
	router.GET("/:shortUrl", s.shortUrlRedirect)
	return router
}
