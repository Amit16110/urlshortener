package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/amit16110/urlshortener/utils"
	"github.com/gin-gonic/gin"
)

type user struct {
	id    uint64
	url   string
	email string
}

var store map[string]user

func hashGenerators() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint64()
}

func (s *Server) urlshortener(c *gin.Context) {
	store = make(map[string]user)
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

	users := user{id: id, email: request.Email, url: request.Url}

	store[shortUrl] = users

	fmt.Println("after", store[shortUrl].email)

	c.JSON(http.StatusOK, gin.H{
		"id":         id,
		"name":       request.Name,
		"short_url":  fmt.Sprintf("http://localhost:8080/%s", shortUrl),
		"created_at": time.Now().UTC(),
	})

}

func (s *Server) shortUrlRedirect(c *gin.Context) {
	sUrl := c.Param("shortUrl")
	val := store[sUrl]

	c.Redirect(302, val.url)
}
