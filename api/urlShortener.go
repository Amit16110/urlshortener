package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/amit16110/urlshortener/utils"
	"github.com/gin-gonic/gin"
)

func hashGenerators() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint64()
}

func (s *Server) urlshortener(c *gin.Context) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Url   string `json:"url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	// hash generator.
	id := hashGenerators()

	shortUrl := utils.Encode(id)

	c.JSON(http.StatusOK, gin.H{
		"id":         id,
		"name":       request.Name,
		"short_url":  fmt.Sprintf("http://localhost:8080/%s", shortUrl),
		"created_at": time.Now().UTC(),
	})

}
